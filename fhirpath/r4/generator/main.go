package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gofhir/generator/base"
	"github.com/iancoleman/strcase"
)

const iPkg = "github.com/gofhir/fhirpath/interpreter"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f := jen.NewFile("r4")

	var typesBundle base.Bundle[base.StructureDefinition]
	var resourcesBundle base.Bundle[base.StructureDefinition]
	profiles, err := os.ReadFile("../../assets/definitions/profiles-types.json")
	check(err)
	check(json.Unmarshal(profiles, &typesBundle))
	profiles, err = os.ReadFile("../../assets/definitions/profiles-resources.json")
	check(err)
	check(json.Unmarshal(profiles, &resourcesBundle))

	var entries []base.BundleEntry[base.StructureDefinition] = append(typesBundle.Entry, resourcesBundle.Entry...)
	var typeAliases map[string][]string = make(map[string][]string)

	for _, sd := range entries {
		s := sd.Resource

		if s.Name == "" {
			log.Fatal("An invalid StructureDefinition snuck in")
		}

		log.Println("Processing", s.Name)

		typeSpecName := strcase.ToLowerCamel(s.Name + "TypeSpec")

		if s.BaseDefinition != "" {
			baseName := strings.TrimPrefix(s.BaseDefinition, "http://hl7.org/fhir/StructureDefinition/")
			baseTypeName := strcase.ToLowerCamel(baseName + "TypeSpec")
			f.Var().Id(typeSpecName).Op("=").Qual(iPkg, "NewTypeSpecWithBase").Call(
				jen.Qual(iPkg, "NewFQTypeName").Call(jen.Lit("FHIR"), jen.Lit(s.Type)),
				jen.Id(baseTypeName),
			)
		} else {
			f.Var().Id(typeSpecName).Op("=").Qual(iPkg, "NewTypeSpec").Call(
				jen.Qual(iPkg, "NewFQTypeName").Call(jen.Lit("FHIR"), jen.Lit(s.Type)),
			)
		}

		for _, snap := range s.Snapshot.Element {
			if len(snap.Type) > 1 {
				path := strings.TrimSuffix(snap.Path, "[x]")
				typeAliases[path] = make([]string, len(snap.Type))
				for i, t := range snap.Type {
					typeAliases[path][i] = strcase.ToCamel(t.Code)
				}
			}
		}

	}

	f.Type().Id("R4Model").Struct()

	selfAccessor := jen.Id("a").Op("*").Id("R4Model")

	f.Func().Params(selfAccessor).Id("AsType").Params(jen.Id("node").Interface(), jen.Id("name").Qual(iPkg, "FQTypeNameAccessor")).Params(jen.Interface(), jen.Error()).Block(
		jen.Panic(jen.Lit("TODO AsType")),
	)

	f.Func().Params(selfAccessor).Id("TypeSpec").Params(jen.Id("node").Interface()).Qual(iPkg, "TypeSpecAccessor").Block(
		jen.Panic(jen.Lit("TODO TypeSpec")),
	)

	f.Func().Params(selfAccessor).Id("Equal").Params(jen.Id("node").Interface(), jen.Id("other").Interface()).Bool().Block(
		jen.Panic(jen.Lit("TODO Equal")),
	)

	f.Func().Params(selfAccessor).Id("Equivalent").Params(jen.Id("node").Interface(), jen.Id("other").Interface()).Bool().Block(
		jen.Panic(jen.Lit("TODO Equivalent")),
	)

	// f.Func().Params(selfAccessor).Id("Navigate").Params(jen.Id("node").Interface(), jen.Id("name").String()).Params(jen.Interface(), jen.Error()).Block(
	// 	generateNavigate(f),
	// )

	f.Func().Params(selfAccessor).Id("Children").Params(jen.Id("node").Interface()).Params(jen.Qual(iPkg, "ColAccessor"), jen.Error()).Block(
		jen.Panic(jen.Lit("TODO Children")),
	)

	// Type aliases
	var values []jen.Code = make([]jen.Code, 0, len(typeAliases))
	for path, types := range typeAliases {
		// values = append(values, jen.Lit(path), jen.Index().String().ValuesFunc(func(group *jen.Group) {
		// 	for _, t := range types {
		// 		group.Lit(t)
		// 	}
		// }))
		values = append(values, jen.Lit(path).Op(":").Index().String().ValuesFunc(func(group *jen.Group) {
			for _, t := range types {
				group.Lit(t)
			}
		},
		))
	}

	f.Var().Id("typeAliases").Op("=").Map(jen.String()).Index().String().Values(
		values...,
	)

	file, err := os.Create("r4.go")
	check(err)
	defer file.Close()
	check(f.Render(file))
}

// func generateNavigate(f *jen.File) jen.Code {
// 	return jen.If(jen.Id("node").Assert(jen.Map(jen.String()).Interface())).Block(
// 		jen.Return(jen.Id("node").Index(jen.Id("name")), jen.Nil()),
// 	).Else().Block(
// 		jen.Panic(jen.Lit("TODO Other node")),
// 	)
// }
