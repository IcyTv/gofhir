package main

import "github.com/dave/jennifer/jen"

// TODO can we parse this from the spec?
var primitiveTypes = map[string]struct{}{
	"base64Binary": {},
	// FIXME I think we need a way to distinguish true/false and unset! Currently we just omit the field if it's false...
	//       the usual way to do this is to use a *boolean...
	"boolean":     {},
	"canonical":   {},
	"code":        {},
	"date":        {},
	"dateTime":    {},
	"decimal":     {},
	"id":          {},
	"instant":     {},
	"integer":     {},
	"integer64":   {},
	"markdown":    {},
	"oid":         {},
	"string":      {},
	"positiveInt": {},
	"time":        {},
	"unsignedInt": {},
	"uri":         {},
	"url":         {},
	"uuid":        {},
}

func IsPrimitiveType(t string) bool {
	_, ok := primitiveTypes[t]
	return ok
}

func (st *StructureTree) GeneratePrimitiveExtensionField() jen.Code {
	return jen.Id(SanitizeFieldName(st.Name) + "PrimitiveExtension").Op("*").Id("PrimitiveExtension").Tag(map[string]string{"bson": st.Name + "_extension,omitempty", "json": "_" + st.Name + ",omitempty"})
}
