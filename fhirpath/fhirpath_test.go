package fhirpath

import (
	"log"
	"testing"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
	"github.com/gofhir/fhirpath/test"
	"github.com/stretchr/testify/assert"
)

func TestSimpleLexing(t *testing.T) {
	expr := "Observation.value as System.String"
	is := antlr.NewInputStream(expr)

	lexer := parser.NewFHIRPathLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		log.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func TestCompileLiteral(t *testing.T) {
	path, err := Compile("true")

	assert.Nil(t, err, "No error should be returned")
	if assert.NotNil(t, path, "Path Expected") {
		assert.NotNil(t, path.evaluator, "Evaluator Expected")
	}
}

func TestCompileEmpty(t *testing.T) {
	path, err := Compile("")

	assert.Nil(t, path, "No path expected")
	if assert.NotNil(t, err, "Error expected") {
		if assert.NotNil(t, err.Items(), "Items expected") {
			assert.Len(t, err.Items(), 1, "One error expected")
		}
	}
}

func TestCompileInvalid(t *testing.T) {
	path, err := Compile("xxx$#@yyy")

	assert.Nil(t, path, "No path expected")
	if assert.NotNil(t, err, "Error expected") {
		if assert.NotNil(t, err.Items(), "Items expected") {
			assert.Len(t, err.Items(), 2)
		}
	}
}

func TestExecuteEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "{}", nil)
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 0, res.Count())
	}
}

func TestExecuteSingle(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "length()", interpreter.NewString("This is a test!"))
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 1, res.Count())
		assert.Equal(t, interpreter.NewInteger(15), res.Get(0))
	}
}

func TestExecuteMulti(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "union('value2' | 'value8')", interpreter.NewString("value5"))
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 3, res.Count())
		assert.Equal(t, interpreter.NewString("value5"), res.Get(0))
		assert.Equal(t, interpreter.NewString("value2"), res.Get(1))
		assert.Equal(t, interpreter.NewString("value8"), res.Get(2))
	}
}

func TestSimpleDate(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "@2019-01-01", nil)
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 1, res.Count())
		assert.Equal(t, interpreter.NewDate(time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.Local)), res.Get(0))
	}
}
