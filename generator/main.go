package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

func main() {
	// Path relative to the project root -> "<project_root>/assets/definitions"
	def_path := "assets/definitions"

	types_file, err := os.ReadFile(def_path + "/profiles-types.json")
	check(err)
	types := base.Bundle[base.StructureDefinition]{}
	check(json.Unmarshal(types_file, &types))

	cycles := DetectCycles(types.Entry)
	_ = cycles

	f := jen.NewFile("generated")

	for _, t := range types.Entry {
		if t.Resource.Name == "Base" {
			continue
		}
		if t.Resource.Kind != "resource" && t.Resource.Kind != "complex-type" {
			continue
		}

		// TODO remove
		if t.Resource.Name != t.Resource.Type {
			fmt.Println("Skipping", t.Resource.Name)
			continue
		}

		tree := NewStructureTree()
		for _, e := range t.Resource.Snapshot.Element {
			tree.AddElement(e)
		}

		// fmt.Printf("%+v\n", tree)
		// spew.Dump(tree)
		res := tree.GenerateCode()
		f.Add(res...)
	}

	out_file, err := os.Create("generated/fhirtypes.go")
	check(err)
	defer out_file.Close()
	f.Render(out_file)
}

func DetectCycles(e []base.BundleEntry[base.StructureDefinition]) []string {
	g := graph.New(graph.StringHash, graph.PreventCycles())

	for _, t := range e {
		_ = g.AddVertex(SanitizeStructName(t.Resource.Name))
		for _, e := range t.Resource.Snapshot.Element {
			for _, ty := range e.Type {
				_ = g.AddVertex(SanitizeStructName(ty.Code))

				err := g.AddEdge(SanitizeStructName(t.Resource.Name), SanitizeStructName(ty.Code), graph.EdgeAttribute("path", e.Path))
				if err != nil {
					panic(err)
				}
			}
		}
	}

	panic("TODO")
}

type StructureTree struct {
	Name  string
	Path  string
	Types []string
	Docs  string

	Min int
	Max string

	Next map[string]*StructureTree
}

func NewStructureTree() *StructureTree {
	return &StructureTree{
		Next: make(map[string]*StructureTree),
	}
}

func (st *StructureTree) AddElement(e base.ElementDefinition) {
	if e.Path == st.Path {
		return
	}

	node := st
	strings.Split(e.Path, ".")
	for _, p := range strings.Split(e.Path, ".") {
		if node.Next[p] == nil {
			node.Next[p] = &StructureTree{
				Name: p,
				Path: strings.TrimPrefix(node.Path+"."+p, "."),
				Next: make(map[string]*StructureTree),
			}
		}
		node = node.Next[p]
	}

	var OutTypes []string
	for _, t := range e.Type {
		OutTypes = append(OutTypes, t.Code)
	}

	node.Types = append(node.Types, OutTypes...)
	node.Docs = e.Definition
	node.Min = e.Min
	node.Max = e.Max
}

func (st *StructureTree) GenerateCode() []jen.Code {
	var Out []jen.Code
	if st.Name == "" {
		for _, n := range st.Next {
			res := n.GenerateCode()
			Out = append(Out, res...)
		}

		return Out
	} else {

		if len(st.Types) > 1 {
			name := SanitizeStructName(st.Path)
			var Fields []jen.Code
			for _, t := range st.Types {
				Fields = append(Fields,
					jen.Id(strcase.ToCamel(ConcatEnumName(st.Name, t))).Id(SanitizeTypes(t)).Tag(map[string]string{"json": ConcatEnumName(st.Name, t)}),
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
			fmt.Println("Generating", st.Path)
			var Fields []jen.Code
			var FieldDefs []jen.Code

			for _, n := range st.Next {
				// Fields = append(Fields, jen.Id(SanitizeStructName(&n.Name)).Add(n.GenerateCode()...))
				gen := n.GenerateCode()
				Fields = append(Fields, gen[0])
				FieldDefs = append(FieldDefs, gen[1:]...)
			}

			if len(strings.Split(st.Path, ".")) > 1 {
				return append(
					[]jen.Code{
						jen.Id(SanitizeFieldName(st.Name)).Id(SanitizeStructName(st.Name)),
						jen.Type().Id(SanitizeStructName(st.Name)).Struct(Fields...),
						jen.Line(),
					},
					FieldDefs...,
				)
			} else {
				return append(
					[]jen.Code{
						jen.Type().Id(SanitizeStructName(st.Name)).Struct(Fields...),
						jen.Line(),
					},
					FieldDefs...,
				)
			}
		} else {
			if len(st.Types) == 0 {
				fmt.Println("No types for", st.Name)
				return []jen.Code{
					jen.Id(SanitizeFieldName(st.Name)).Id("interface{}"),
				}
			}

			// return []jen.Code{
			// 	jen.Id(SanitizeTypes(st.Types[0])),
			// }
			return []jen.Code{
				jen.Id(SanitizeFieldName(st.Name)).Id(SanitizeTypes(st.Types[0])).Tag(map[string]string{"json": st.Name}).Comment(st.Docs),
			}
		}

	}
}

func SanitizeStructName(s string) string {
	// Mainly replace "." with "_"
	repl := strings.NewReplacer(".", "_")
	return strcase.ToCamel(repl.Replace(s))
}

func SanitizeTypes(s string) string {
	var typeMapper = map[string]string{
		"String":       "string",
		"Boolean":      "bool",
		"PositiveInt":  "int",
		"UnsignedInt":  "int",
		"Integer":      "int",
		"Integer64":    "int64",
		"Base64Binary": "string",
		"Canonical":    "string",
		"Code":         "string",
		"Date":         "string",
		"DateTime":     "string",
		"Decimal":      "float64",
		"Id":           "string",
		"Instant":      "string",
		"Markdown":     "string",
		"Oid":          "string",
		"Uri":          "string",
		"Url":          "string",
		"Time":         "string",
		"Uuid":         "string",
		"Xhtml":        "string",
	}

	repl := strings.NewReplacer("http://hl7.org/fhirpath/System.", "", "type", "Type")
	ret := strcase.ToCamel(repl.Replace(s))
	if typeMapper[ret] != "" {
		return typeMapper[ret]
	} else {
		return ret
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
