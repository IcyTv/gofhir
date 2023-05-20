package expression

import "github.com/gofhir/fhirpath/interpreter"

type EmptyLiteral struct{}

var emptyLiteral = &EmptyLiteral{}

func NewEmptyLiteral() interpreter.Evaluator {
	return emptyLiteral
}

func (e *EmptyLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return nil, nil
}
