package fhirpath

//go:generate antlr4 -Dlanguage=Go -o parser -package parser -no-listener -visitor FHIRPath.g4

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/errors"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
	"github.com/gofhir/fhirpath/visitors"
	"go.mongodb.org/mongo-driver/bson"
)

func FhirpathToMongoQueryAndProjection(fhirpath string) (bson.M, bson.M) {
	panic("TODO")
}

type Path struct {
	evaluator expression.CollectionExpression
}

func Compile(pathString string) (*Path, *errors.Error) {
	errorColl := visitors.NewErrorItemCollection()
	errorListener := visitors.NewErrorListener(errorColl)

	is := antlr.NewInputStream(pathString)
	lexer := parser.NewFHIRPathLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFHIRPathParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	v := visitors.NewVisitor(errorColl)
	res := p.Expression().Accept(v)

	if errorColl.HasErrors() {
		return nil, errors.NewError("Error when parsing expression", errorColl.Items())
	}

	if !errorColl.HasErrors() && res == nil {
		log.Fatal("No result returned from visitor")
	}

	return &Path{expression.NewCollectionExpression(res.(interpreter.Evaluator))}, nil
}

func Execute(ctx interpreter.ContextAccessor, pathString string, node interface{}) (interpreter.ColAccessor, *errors.Error) {
	path, err := Compile(pathString)
	if err != nil {
		return nil, err
	}

	return path.Execute(ctx, node)
}

func (p *Path) Execute(ctx interpreter.ContextAccessor, node interface{}) (interpreter.ColAccessor, *errors.Error) {
	res, err := p.evaluator.Evaluate(ctx, node, nil)
	if err != nil {
		return nil, errors.NewError(err.Error(), nil)
	}
	return res.(interpreter.ColAccessor), nil
}
