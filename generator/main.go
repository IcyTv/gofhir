package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/dave/jennifer/jen"
	"github.com/dominikbraun/graph"
	"github.com/gofhir/generator/base"
	"github.com/iancoleman/strcase"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var wg sync.WaitGroup

func main() {
	// Path relative to the project root -> "<project_root>/assets/definitions"
	def_path := "assets/definitions"
	out_path := "generated"

	GenerateFile(def_path+"/profiles-types.json", out_path+"/fhirtypes.go", out_path+"/httptypes.go")
	GenerateFile(def_path+"/profiles-resources.json", out_path+"/fhirresources.go", out_path+"/httpresources.go")
	GenerateFile(def_path+"/profiles-others.json", out_path+"/fhirothers.go", out_path+"/httpothers.go")

	wg.Wait()
}

func GenerateFile(infile string, outfile string, http_outfile string) {
	types_file, err := os.ReadFile(infile)
	check(err)

	// TODO we now iterate over the entries multiple times
	types := base.Bundle[base.StructureDefinition]{}
	check(json.Unmarshal(types_file, &types))
	capabilites := base.Bundle[base.CapabilityStatement]{}
	check(json.Unmarshal(types_file, &capabilites))

	var entries []base.StructureDefinition
	var capabilityStatements []base.CapabilityStatement
	for _, e := range types.Entry {
		entries = append(entries, e.Resource)
	}
	for _, e := range capabilites.Entry {
		capabilityStatements = append(capabilityStatements, e.Resource)
	}

	go GenerateStructures(entries, outfile)
	go GenerateHttp(capabilityStatements, http_outfile)
}

func GenerateStructures(entries []base.StructureDefinition, outfile string) {
	wg.Add(1)
	fmt.Println("Generating", outfile)

	cycles := DetectCycles(entries)

	f := jen.NewFile("generated")

	for _, t := range entries {
		if t.Name == "Base" {
			continue
		}
		if t.Kind != "resource" && t.Kind != "complex-type" {
			continue
		}

		// TODO remove
		if t.Name != t.Type {
			continue
		}

		tree := NewStructureTree()
		for _, e := range t.Snapshot.Element {
			tree.AddElement(e)
		}

		res := tree.GenerateCode(cycles)
		f.Add(res...)
	}

	out_file, err := os.Create(outfile)
	check(err)
	defer out_file.Close()
	check(f.Render(out_file))

	wg.Done()
}

func GenerateHttp(caps []base.CapabilityStatement, outfile string) {
	wg.Add(1)
	fmt.Println("Generating", outfile)

	f := jen.NewFile("generated")

	f.ImportAlias("github.com/json-iterator/go", "jsoniter")

	var routes []jen.Code

	for _, cap := range caps {
		for _, rest := range cap.Rest {
			// f.Func().Id(rest.Resource + "Group").Params().Return(jen.Qual("github.com/gin-gonic/gin", "RouterGroup")).Block(
			// 	jen.Panic(jen.Lit("Not implemented")),
			// )

			for _, resource := range rest.Resource {
				inner := []jen.Code{
					jen.Id("g").Op(":=").Id("r").Dot("Group").Call(jen.Lit("/" + resource.Type)),
				}

				for _, interaction := range resource.Interaction {
					switch interaction.Code {
					case "read":
						inner = append(inner, GenerateRead(resource.Type))
					case "update":
						inner = append(inner, GenerateUpdate(resource.Type))
					case "create":
						inner = append(inner, GenerateCreate(resource.Type))
					default:
						fmt.Println("Unknown interaction", interaction.Code)
					}

				}

				inner = append(inner, jen.Return(jen.Id("g")))

				f.Func().Id(resource.Type+"Routes").Params(
					jen.Id("r").Op("*").Qual("github.com/gin-gonic/gin", "Engine"),
				).Op("*").Qual("github.com/gin-gonic/gin", "RouterGroup").Block(
					inner...,
				)

				routes = append(routes, jen.Id(resource.Type+"Routes").Call(jen.Id("r")))
			}
		}
	}

	if len(routes) > 0 {
		//TODO this leads to name collisisons if we have multiple files...
		f.Func().Id("Routes").Params(
			jen.Id("r").Op("*").Qual("github.com/gin-gonic/gin", "Engine"),
		).Block(
			routes...,
		)
	}

	out_file, err := os.Create(outfile)
	check(err)
	defer out_file.Close()
	check(f.Render(out_file))

	wg.Done()
}

func DetectCycles(e []base.StructureDefinition) map[string]struct{} {
	Out := make(map[string]struct{})
	g := graph.New(graph.StringHash, graph.PreventCycles())

	for _, t := range e {
		resourceName := SanitizeStructName(t.Name)
		_ = g.AddVertex(resourceName)
		for _, e := range t.Snapshot.Element {
			for _, ty := range e.Type {
				typeName := SanitizeStructName(ty.Code)
				_ = g.AddVertex(typeName)

				createsCycle, err := graph.CreatesCycle(g, resourceName, typeName)
				check(err)

				if createsCycle {
					Out[e.Path] = struct{}{}
				} else {
					err := g.AddEdge(resourceName, typeName)
					check(err)
				}
			}
		}
	}

	return Out
}

type StructureTree struct {
	Name  string
	Path  string
	Types []string
	Docs  string

	Min int
	Max string

	IsResource bool

	Next map[string]*StructureTree
}

func NewStructureTree() *StructureTree {
	return &StructureTree{
		Next: make(map[string]*StructureTree),
	}
}

func (st *StructureTree) AddElement(e base.ElementDefinition) {
	node := st
	strings.Split(e.Path, ".")
	for idx, p := range strings.Split(e.Path, ".") {
		if node.Next[p] == nil {
			node.Next[p] = &StructureTree{
				Name: p,
				Path: strings.TrimPrefix(node.Path+"."+p, "."),
				Next: make(map[string]*StructureTree),
			}
		}
		node = node.Next[p]
		if idx == 0 {
			node.IsResource = true
		}
	}

	var OutTypes []string
	for _, t := range e.Type {
		if len(t.Extension) > 0 && t.Extension[0].ValueUrl != "" {
			OutTypes = append(OutTypes, t.Extension[0].ValueUrl)
		} else {
			OutTypes = append(OutTypes, t.Code)
		}
	}

	node.Types = append(node.Types, OutTypes...)
	node.Docs = e.Definition
	node.Min = e.Min
	node.Max = e.Max
}

func (st *StructureTree) GenerateCode(cycles map[string]struct{}) []jen.Code {
	var Out []jen.Code
	if st.Name == "" {
		for _, n := range st.Next {
			res := n.GenerateCode(cycles)
			Out = append(Out, res...)
		}

		return Out
	} else {
		if len(st.Types) > 1 {
			name := SanitizeStructName(st.Path)
			var Fields []jen.Code
			for _, t := range st.Types {
				tags := map[string]string{
					"json": ConcatEnumName(st.Name, t) + ",omitempty",
					// "mapstructure": ConcatEnumName(st.Name, t),
					"bson": ",omitempty",
				}

				Fields = append(Fields,
					jen.Id(strcase.ToCamel(ConcatEnumName(st.Name, t))).Add(SanitizeTypes(t)).Tag(tags),
				)
			}

			str := jen.Type().Id(name).Struct(
				Fields...,
			)

			return append(
				[]jen.Code{
					jen.Id(SanitizeStructName(st.Path)),
					str,
					jen.Line(),
				},
				Out...,
			)
		} else if len(st.Next) > 0 {
			var Fields []jen.Code
			var FieldDefs []jen.Code
			var UnmarshalFields []jen.Code

			for _, n := range st.Next {
				gen := n.GenerateCode(cycles)
				Fields = append(Fields, gen[0])
				FieldDefs = append(FieldDefs, gen[1:]...)
				UnmarshalFields = append(UnmarshalFields, GenerateUnmarshal(SanitizeFieldName(n.Name), n))
			}

			UnmarshalBlock := []jen.Code{
				jen.Var().Id("asMap").Map(jen.String()).Qual("github.com/json-iterator/go", "RawMessage"),
				// Unmarshal the JSON into a map
				jen.If(
					jen.Err().Op(":=").Qual("github.com/json-iterator/go", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("asMap")),
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Return(jen.Err()),
				),
			}
			if st.IsResource {
				// If this is a resource, we need to check for the resourceType field and assert that it's the correct type
				// FIXME this is wrong and needs to be fixed.
				//       generally the "resourceType" should only appear on json base types...
				//       currently we make the resourceType field optional
				UnmarshalBlock = append(UnmarshalBlock,
					jen.If(
						// jen.String().Params(jen.Id("asMap").Index(jen.Lit("resourceType"))).Op("!=").Lit("\""+st.Name+"\""),
						jen.List(jen.Id("resource"), jen.Id("ok")).Op(":=").Id("asMap").Index(jen.Lit("resourceType")),
						jen.Id("ok").Op("&&").String().Call(jen.Id("resource")).Op("!=").Lit("\""+st.Name+"\""),
					).Block(
						jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("resourceType is not %s"), jen.Lit(st.Name))),
					),
				)
			}

			UnmarshalBlock = append(UnmarshalBlock, UnmarshalFields...)
			UnmarshalBlock = append(UnmarshalBlock,
				jen.Return(jen.Nil()),
			)

			unmarshalJson := jen.Func().Params(jen.Id("out").Op("*").Id(SanitizeStructName(st.Path))).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Params(jen.Error()).Block(
				UnmarshalBlock...,
			).Line()

			if len(strings.Split(st.Path, ".")) > 1 {
				return append(
					[]jen.Code{
						jen.Id(SanitizeFieldName(st.Name)).Op("*").Id(SanitizeStructName(st.Path)).Tag(map[string]string{"binding": "omitempty", "bson": ",omitempty"}),
						unmarshalJson,
						jen.Type().Id(SanitizeStructName(st.Path)).Struct(Fields...),
						jen.Line(),
					},
					FieldDefs...,
				)
			} else {
				//FIXME move this logic to the if statement at the beginning?
				return append(
					[]jen.Code{
						unmarshalJson,
						jen.Type().Id(SanitizeStructName(st.Name)).Struct(
							append(Fields,
								jen.Id("ResourceType").String().Tag(map[string]string{"json": "resourceType", "binding": "omitempty", "bson": "-"}),
							)...,
						),
						jen.Line(),
					},
					FieldDefs...,
				)
			}
		} else {
			if len(st.Types) == 0 {
				return []jen.Code{
					jen.Id(SanitizeFieldName(st.Name)).Id("interface{}"),
				}
			}

			fieldType := SanitizeTypes(st.Types[0])

			if _, ok := cycles[st.Path]; ok {
				// fieldType = "*" + fieldType
				fieldType = jen.Op("*").Add(fieldType)
			}

			if st.Max == "*" {
				fieldType = jen.Index().Add(fieldType)
			}

			tag := map[string]string{
				"json": st.Name + ",omitempty",
				// "mapstructure": st.Name,
				"bson": ",omitempty",
			}
			if st.Min > 0 {
				tag["binding"] = "required"
			}

			if st.Name == "id" {
				tag["json"] = "id,omitempty"
				tag["bson"] = "_id,omitempty"
			}

			return []jen.Code{
				jen.Id(SanitizeFieldName(st.Name)).Add(fieldType).Tag(tag).Comment(st.Docs),
			}
		}

	}
}

func GenerateUnmarshal(fieldName string, n *StructureTree) jen.Code {

	if strings.HasSuffix(n.Name, "[x]") {
		// Try unmarshalling into each of the possible types
		// Say we have the following:
		// n.Name = "value[x]"
		// n.Types = ["string", "integer"]
		// We want to generate the following:
		//   1. Try unmarshalling asMap["valueString"] into out.valueString
		//   2. Try unmarshalling asMap["valueInteger"] into out.valueInteger
		//   3. Return an error if neither of those work and the field is required

		var UnmarshalFields []jen.Code
		for _, t := range n.Types {
			fieldName := ConcatEnumName(n.Name, t)
			Unmarshal := jen.If(
				jen.Id("err").Op(":=").Qual("github.com/json-iterator/go", "Unmarshal").Call(jen.Id("asMap").Index(jen.Lit(fieldName)), jen.Op("&").Id("out").Dot(strcase.ToCamel(fieldName))),
				jen.Id("err").Op("==").Nil(),
			).Block().Else()

			UnmarshalFields = append(UnmarshalFields, Unmarshal)
		}

		UnmarshalError := jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("could not unmarshal %s into any of the possible types"), jen.Lit(n.Name)))
		if n.Min == 0 {
			UnmarshalError = jen.Line()
		}

		Fields := jen.Add(UnmarshalFields...).Block(
			UnmarshalError,
		)

		return Fields
	} else {
		// Unmarshal this straight into the field
		Unmarshal := jen.If(
			jen.Id("err").Op(":=").Qual("github.com/json-iterator/go", "Unmarshal").Call(jen.Id("asMap").Index(jen.Lit(n.Name)), jen.Op("&").Id("out").Dot(fieldName)),
			jen.Id("err").Op("!=").Nil(),
		).Block(
			jen.Return(jen.Id("err")),
		).Line()

		if n.Min == 0 {
			// If the field is optional test if the len(asMap[fieldName]) > 0
			Unmarshal = jen.If(
				jen.Len(jen.Id("asMap").Index(jen.Lit(n.Name))).Op(">").Lit(0),
			).Block(
				Unmarshal,
			)
		}

		return Unmarshal
	}
}

func SanitizeStructName(s string) string {
	// Mainly replace "." with "_"
	repl := strings.NewReplacer(".", "_")
	return strcase.ToCamel(repl.Replace(s))
}

func SanitizeTypes(s string) jen.Code {
	var typeMapper = map[string]jen.Code{
		"String":       jen.String(),
		"Boolean":      jen.Bool(),
		"PositiveInt":  jen.Int(),
		"UnsignedInt":  jen.Int(),
		"Integer":      jen.Int(),
		"Integer64":    jen.Int64(),
		"Base64Binary": jen.Id("Base64Binary"),
		"Canonical":    jen.String(),
		"Code":         jen.String(),
		"Decimal":      jen.Float64(),
		"Id":           jen.Op("*").Qual("go.mongodb.org/mongo-driver/bson/primitive", "ObjectID"),
		"Instant":      jen.Qual("time", "Time"),
		"Markdown":     jen.String(),
		"Oid":          jen.String(),
		"Uri":          jen.String(),
		"Url":          jen.Qual("net/url", "URL"),
		"Time":         jen.Id("Time"),
		"Date":         jen.Id("Date"),
		"DateTime":     jen.Id("DateTime"),
		"Uuid":         jen.Qual("github.com/google/uuid", "UUID"),
		"Xhtml":        jen.String(),
	}

	repl := strings.NewReplacer("http://hl7.org/fhirpath/System.", "", "type", "Type")
	ret := strcase.ToCamel(repl.Replace(s))
	// if typeMapper[ret] != "" {
	// 	return typeMapper[ret]
	// } else {
	// 	return ret
	// }
	if v, ok := typeMapper[ret]; ok {
		return v
	} else {
		return jen.Id(ret)
	}
}

func SanitizeFieldName(s string) string {
	repl := strings.NewReplacer("type", "Type", "[x]", "", "map", "Map")
	return strcase.ToCamel(repl.Replace(s))
}

func ConcatEnumName(s string, ty string) string {
	return strings.Replace(
		s,
		"[x]",
		strcase.ToCamel(ty),
		1,
	)
}
