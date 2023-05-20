package main

import (
	. "github.com/dave/jennifer/jen"
)

const gin = "github.com/gin-gonic/gin"
const primitive = "go.mongodb.org/mongo-driver/bson/primitive"
const bson = "go.mongodb.org/mongo-driver/bson"
const db = "github.com/gofhir/db"

func coll(resource string) Code {
	return Id("coll").Op(":=").Qual(db, "FhirDatabase").Dot("Collection").Call(Lit(resource))
}

func operation(method string, path string, code ...Code) Code {
	return Id("g").Dot(method).Call(
		Lit(path),
		Func().Params(
			Id("c").Op("*").Qual(gin, "Context"),
		).Block(code...),
	)
}

func retErr(err string) Code {
	return If(Id("err").Op("!=").Nil()).Block(
		Id("c").Dot("JSON").Call(Qual("net/http", err), Qual(gin, "H").Values(
			Lit("error").Op(":").Id("err").Dot("Error").Call().Op(","),
		)),
		Qual("fmt", "Println").Call(Id("err")),
		Return(),
	)
}

func toOid(name string) Code {
	return List(Id("objId"), Id("err")).Op(":=").Qual(primitive, "ObjectIDFromHex").Call(Id("c").Dot("Param").Call(Lit(name))).Line().Add(retErr("StatusBadRequest"))
}

func GenerateRead(resource string) Code {
	// TODO this is terribly unreadable. Find a better way...
	return operation("GET", "/:id",
		coll(resource),
		toOid("id"),
		Id("res").Op(":=").Id("coll").Dot("FindOne").Call(Qual("context", "TODO").Call(), Qual(bson, "M").Values(
			Lit("_id").Op(":").Id("objId"),
		)),
		Id("result").Op(":=").Id(resource).Values(),
		Id("err").Op("=").Id("res").Dot("Decode").Call(Op("&").Id("result")),
		retErr("StatusBadRequest"),
		Id("format").Op(":=").Id("c").Dot("DefaultQuery").Call(Lit("_format"), Lit("json")),
		Println(Id("format")),
		If(
			Id("format").Op("==").Lit("json"),
		).Block(
			// Add the field "resourceType" to the result
			Id("result").Dot("ResourceType").Op("=").Lit(resource),
			Id("c").Dot("JSON").Call(Qual("net/http", "StatusOK"), Id("result")),
		).Else().If(Id("format").Op("==").Lit("xml")).Block(
			Id("c").Dot("XML").Call(Qual("net/http", "StatusOK"), Id("result")),
		).Else().Block(
			Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Values(
				Lit("error").Op(":").Lit("Invalid _format parameter"),
			)),
		),
	)
}

// # Rejecting Updates
// 400 Bad Request - resource could not be parsed or failed basic FHIR validation rules (or multiple matches were found for conditional criteria)
// 401 Unauthorized - authorization is required for the interaction that was attempted
// 404 Not Found - resource type not supported, or not a FHIR end-point
// 405 Method Not Allowed - the resource did not exist prior to the update, and the server does not allow client defined ids
// 409 Conflict/412 Precondition Failed - version conflict management - see below
func GenerateUpdate(resource string) Code {
	return Id("g").Dot("PUT").Call(
		Lit("/:id"),
		Func().Params(
			Id("c").Op("*").Qual(gin, "Context"),
		).Block(
			Id("a").Op(":=").Id(resource).Block(),
			// FIXME remove the need for correct ObjectId deserialization here, since we should ignore the id anyways
			If(
				Id("errA").Op(":=").Id("c").Dot("ShouldBind").Call(Op("&").Id("a")),
				Id("errA").Op("!=").Nil(),
			).Block(
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Values(
					Lit("error").Op(":").Id("errA").Dot("Error").Call(),
				)),
				Qual("fmt", "Println").Call(Id("errA")),
				Return(),
			),

			coll(resource),
			toOid("id"),
			List(Id("result"), Id("err")).Op(":=").Id("coll").Dot("ReplaceOne").Call(Qual("context", "TODO").Call(), Qual(bson, "M").Values(
				Lit("_id").Op(":").Id("objId"),
			), Id("a")),
			If(
				Id("err").Op("!=").Nil(),
			).Block(
				// If the resource does not exist, return 405, method not allowed, since we don't allow client defined ids
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusMethodNotAllowed"), Qual(gin, "H").Values(
					Lit("error").Op(":").Lit("Resource does not exist"),
				)),
			),
			Qual("fmt", "Println").Call(Id("result")),

			Id("c").Dot("String").Call(Qual("net/http", "StatusOK"), Lit("Success")),
		),
	)
}

func GenerateCreate(resource string) Code {
	return Id("g").Dot("POST").Call(
		Lit(""),
		Func().Params(
			Id("c").Op("*").Qual(gin, "Context"),
		).Block(
			Id("a").Op(":=").Id(resource).Block(),
			// FIXME remove the need for correct ObjectId deserialization here, since we should ignore the id anyways
			If(
				Id("errA").Op(":=").Id("c").Dot("ShouldBind").Call(Op("&").Id("a")),
				Id("errA").Op("!=").Nil(),
			).Block(
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Block(
					Lit("error").Op(":").Id("errA").Dot("Error").Call().Op(","),
				)),
				Qual("fmt", "Println").Call(Id("errA")),
				Return(),
			),

			Id("objId").Op(":=").Qual(primitive, "NewObjectID").Call(),
			Id("a").Dot("Id").Op("=").Op("&").Id("objId"),

			coll(resource),
			List(Id("result"), Id("err")).Op(":=").Id("coll").Dot("InsertOne").Call(Qual("context", "TODO").Call(), Id("a")),
			retErr("StatusBadRequest"),
			Qual("fmt", "Println").Call(Id("result")),

			Id("c").Dot("String").Call(Qual("net/http", "StatusOK"), Lit("Success")),
		),
	)
}
