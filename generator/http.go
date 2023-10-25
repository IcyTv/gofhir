package main

import (
	"github.com/dave/jennifer/jen"
	. "github.com/dave/jennifer/jen"
)

const gin = "github.com/gin-gonic/gin"
const primitive = "go.mongodb.org/mongo-driver/bson/primitive"
const bson = "go.mongodb.org/mongo-driver/bson"
const options = "go.mongodb.org/mongo-driver/mongo/options"
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

		Var().Id("headers").Id("FHIRHeader"),
		Id("c").Dot("BindHeader").Call(Op("&").Id("headers")),
		Qual("fmt", "Printf").Call(Lit("%+v\n"), Id("headers")),

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

			If(
				Id("result").Dot("MatchedCount").Op("==").Lit(0),
			).Block(
				//If no id element is provided, or the id disagrees with the id in the URL, the server SHALL respond with an HTTP 400 Bad Request error code
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Values(
					Lit("error").Op(":").Lit("Id in body does not match id in URL"),
				)),
			),

			Qual("fmt", "Println").Call(Id("result")),

			// Id("c").Dot("String").Call(Qual("net/http", "StatusOK"), Lit("Success")),
			// TODO what do we return here?
			Id("c").Dot("Header").Call(jen.Lit("Location"), jen.Lit("/"+resource+"/").Op("+").Id("objId").Dot("Hex").Call().Op("+").Lit("/history/").Op("+").Lit("TODO")),
			Id("c").Dot("JSON").Call(Qual("net/http", "StatusOK"), Map(String()).Interface().Values()),
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

			//TODO getting the time on our local machine is a shit way to do this...
			Id("currentTime").Op(":=").Qual("time", "Now").Call(),

			Id("a").Dot("Meta").Op("=").Op("&").Id("Meta").Values(
				Id("VersionId").Op(":").Lit(1),
				Id("LastUpdated").Op(":").Op("&").Id("currentTime"),
			),

			coll(resource),
			List(Id("result"), Id("err")).Op(":=").Id("coll").Dot("InsertOne").Call(Qual("context", "TODO").Call(), Id("a")),
			retErr("StatusBadRequest"),
			Qual("fmt", "Println").Call(Id("result")),

			Id("c").Dot("Header").Call(jen.Lit("Location"), jen.Lit("/"+resource+"/").Op("+").Id("objId").Dot("Hex").Call()),

			Id("c").Dot("JSON").Call(Qual("net/http", "StatusCreated"), Id("a")),
		),
	)
}

func GeneratePatch(resource string) Code {
	return Id("g").Dot("PATCH").Call(
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

			Id("update").Op(":=").Map(String()).Map(String()).Interface().Values(
				Lit("$set").Op(":").Values(),
				Lit("$unset").Op(":").Values(),
				Lit("$rename").Op(":").Values(),
			),

			// 1. If mime-type is application/json-patch+json, convert to MongoDB patch format
			If(
				Id("c").Dot("GetHeader").Call(Lit("Content-Type")).Op("==").Lit("application/json-patch+json"),
			).Block(
				Var().Id("body").Index().Byte(),
				Id("c").Dot("Request").Dot("Body").Dot("Read").Call(Id("body")),
				Id("patch").Op(",").Id("err").Op(":=").Qual("github.com/evanphx/json-patch/v5", "DecodePatch").Call(
					Id("body"),
				),
				retErr("StatusBadRequest"),

				// TODO we should outsource all this patch stuff to a separate function, and then just call it to convert from the
				// json-patch to a mongo update
				For(
					List(Id("_"), Id("op")).Op(":=").Range().Id("patch"),
				).Block(
					Id("path").Op(",").Id("err").Op(":=").Id("op").Dot("Path").Call(),
					retErr("StatusBadRequest"),

					Id("path").Op("=").Qual("strings", "ReplaceAll").Call(Id("path"), Lit("/"), Lit(".")),

					If(Id("op").Dot("Kind").Call().Op("==").Lit("add")).Block(
						Id("value").Op(",").Id("err").Op(":=").Id("op").Dot("ValueInterface").Call(),
						retErr("StatusBadRequest"),
						//TODO panic if we try to add a field that already exists
						Id("update").Index(Lit("$set")).Index(Id("path")).Op("=").Id("value"),
					).Else().If(Id("op").Dot("Kind").Call().Op("==").Lit("remove")).Block(
						Id("update").Index(Lit("$unset")).Index(Id("path")).Op("=").Lit(""),
					).Else().If(Id("op").Dot("Kind").Call().Op("==").Lit("replace")).Block(
						// Panic(Lit("TODO replace")),
						// TODO we need to check if the field exists... Right now it's pretty much the same as add
						Id("value").Op(",").Id("err").Op(":=").Id("op").Dot("ValueInterface").Call(),
						retErr("StatusBadRequest"),
						Id("update").Index(Lit("$set")).Index(Id("path")).Op("=").Id("value"),
					).Else().If(Id("op").Dot("Kind").Call().Op("==").Lit("move")).Block(
						List(Id("from"), Id("err")).Op(":=").Id("op").Dot("From").Call(),
						retErr("StatusBadRequest"),
						List(Id("to"), Id("err")).Op(":=").Id("op").Dot("Path").Call(),
						retErr("StatusBadRequest"),

						Id("from").Op("=").Qual("strings", "ReplaceAll").Call(Id("from"), Lit("/"), Lit(".")),
						Id("to").Op("=").Qual("strings", "ReplaceAll").Call(Id("to"), Lit("/"), Lit(".")),

						Id("update").Index(Lit("$rename")).Index(Id("from")).Op("=").Id("to"),
					).Else().If(Id("op").Dot("Kind").Call().Op("==").Lit("copy")).Block(
						Panic(Lit("TODO copy")),
					).Else().If(Id("op").Dot("Kind").Call().Op("==").Lit("test")).Block(
						Panic(Lit("TODO test")),
					).Else().Block(
						Id("err").Op("=").Qual("errors", "New").Call(Lit("Unsupported operation")),
						retErr("StatusBadRequest"),
					),
				),
			).Else().Block(
				Panic(Lit("Unsupported content type (for now)")),
			),

			List(Id("res"), Id("err")).Op(":=").Id("coll").Dot("UpdateByID").Call(Qual("context", "TODO").Call(), Id("objId"), Id("update")),
			retErr("StatusBadRequest"),
			Qual("fmt", "Println").Call(Id("res")),
		),
	)
}

func GenerateVRead(resource string) Code {
	return Id("g").Dot("GET").Call(
		Lit("/:id/_history/:vid"),
		Func().Params(
			Id("c").Op("*").Qual(gin, "Context"),
		).Block(
			coll(resource),
			toOid("id"),
			List(Id("vid64"), Id("err")).Op(":=").Qual("strconv", "ParseUint").Call(Id("c").Dot("Param").Call(Lit("vid")), Lit(10), Lit(32)),
			retErr("StatusBadRequest"),
			Id("vid").Op(":=").Int().Call(Id("vid64")),

			// 1. get the base resource
			Id("res").Op(":=").Id("coll").Dot("FindOne").Call(Qual("context", "TODO").Call(), Qual(bson, "M").Values(
				Lit("_id").Op(":").Id("objId"),
			)),

			// 2. Decode the `SingleResult` struct into a map[string]interface{}
			Id("result").Op(":=").Id(resource).Values(),
			Id("err").Op("=").Id("res").Dot("Decode").Call(Op("&").Id("result")),
			retErr("StatusBadRequest"),

			// 3.a. If the versionid matches, return the resource
			If(
				Id("result").Dot("Meta").Dot("VersionId").Op("==").Id("vid"),
			).Block(
				// Add the field "resourceType" to the result
				Id("result").Dot("ResourceType").Op("=").Lit(resource),
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusOK"), Id("result")),
				Return(),
			),

			// 3.b. If the requested versionId is larger than the current versionId, return 404
			If(
				Id("result").Dot("Meta").Dot("VersionId").Op("<").Id("vid"),
			).Block(
				Id("c").Dot("JSON").Call(Qual("net/http", "StatusNotFound"), Qual(gin, "H").Values(
					Lit("error").Op(":").Lit("Version not found"),
				)),
				Return(),
			),

			// 3.c. If the versionid does not match, get all the deltas for this resource from the `VersioningDatabase`, with a versionid greater than the requested versionid
			// TODO it might be more storage efficient if we don't store the versionId and instead just count the number of deltas...
			List(Id("cursor"), Id("err")).Op(":=").Qual("github.com/gofhir/db", "VersioningDatabase").Dot("Collection").Call(Lit(resource)).Dot("Find").Call(
				Qual("context", "TODO").Call(),
				Qual(bson, "M").Values(
					Lit("resourceId").Op(":").Id("objId"),
					Lit("versionId").Op(":").Qual(bson, "M").Values(
						Lit("$gt").Op(":").Id("vid").Op("-").Lit(1),
					),
				),
				Op("&").Qual(options, "FindOptions").Values(
					Id("Projection").Op(":").Qual(bson, "M").Values(
						Lit("delta").Op(":").Lit(1),
					),
					Id("Sort").Op(":").Qual(bson, "M").Values(
						Lit("versionId").Op(":").Lit(-1),
					),
				),
			),
			retErr("StatusBadRequest"),
			Defer().Id("cursor").Dot("Close").Call(Qual("context", "TODO").Call()),

			// 4. Apply the deltas to the base resource
			For(
				Id("cursor").Dot("Next").Call(Qual("context", "TODO").Call()),
			).Block(
				Qual("fmt", "Println").Call(Lit("Applying delta")),
				// Var().Id("delta").Qual("github.com/r3labs/diff/v3", "Changelog")
				// Var().Id("doc").Map(String()).Qual("github.com/r3labs/diff/v3", "Changelog"),
				Var().Id("doc").Map(String()).Interface(),
				If(
					Id("err").Op(":=").Id("cursor").Dot("Decode").Call(Op("&").Id("doc")),
					Id("err").Op("!=").Nil(),
				).Block(
					Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Values(
						Lit("error").Op(":").Id("err").Dot("Error").Call(),
					)),
					Qual("fmt", "Println").Call(Id("err")),
					Return(),
				),
				Qual("fmt", "Printf").Call(Lit("Applying delta %+v\n"), Id("doc")),
				// Id("delta").Op(":=").Id("doc").Index(Lit("delta")).Op(".").Parens(Qual("github.com/r3labs/diff/v3", "Changelog")),
				// TODO this feels insanely hacky...
				List(Id("marshaled"), Id("err")).Op(":=").Qual("encoding/json", "Marshal").Call(Id("doc").Index(Lit("delta"))),
				retErr("StatusInternalServerError"),
				Id("delta").Op(":=").Qual("github.com/r3labs/diff/v3", "Changelog").Values(),
				Id("err").Op("=").Qual("encoding/json", "Unmarshal").Call(Id("marshaled"), Op("&").Id("delta")),
				retErr("StatusInternalServerError"),

				Qual("fmt", "Println").Call(Id("delta")),
				Id("patchlog").Op(":=").Qual("github.com/r3labs/diff/v3", "Patch").Call(Id("delta"), Op("&").Id("result")),
				If(Id("patchlog").Dot("HasErrors").Call()).Block(
					Id("c").Dot("JSON").Call(Qual("net/http", "StatusBadRequest"), Qual(gin, "H").Values(
						Lit("error").Op(":").Lit("Error generating version"),
					)),
					Qual("fmt", "Printf").Call(Lit("Error applying patches... %d errors\n"), Id("patchlog").Dot("ErrorCount").Call()),

					For(
						List(Id("_"), Id("patch")).Op(":=").Range().Id("patchlog"),
					).Block(
						Qual("fmt", "Printf").Call(Lit("Error: %+v\n"), Id("patch").Dot("Errors")),
					),

					Return(),
				),
			),

			//5. Return the patched resource, making sure to set the correct versionId
			// TODO what about the lastUpdated field? etc?

			// Add the field "resourceType" to the result
			Id("result").Dot("ResourceType").Op("=").Lit(resource),
			// Id("result").Dot("Meta").Dot("VersionId").Op("=").Id("vid"),
			Id("c").Dot("JSON").Call(Qual("net/http", "StatusOK"), Id("result")),
		),
	)
}
