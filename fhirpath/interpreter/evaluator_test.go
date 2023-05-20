package interpreter_test

import (
	"github.com/gofhir/fhirpath/interpreter"
)

type testEvaluator struct {
}

func newTestEvaluator() interpreter.Evaluator {
	return &testEvaluator{}
}

func (t testEvaluator) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return nil, nil
}
