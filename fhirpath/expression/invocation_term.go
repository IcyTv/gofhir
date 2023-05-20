package expression

import "github.com/gofhir/fhirpath/interpreter"

type InvocationTerm struct {
	evaluator interpreter.Evaluator
}

func NewInvocationTerm(evaluator interpreter.Evaluator) *InvocationTerm {
	return &InvocationTerm{evaluator}
}

func (e *InvocationTerm) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	return e.evaluator.Evaluate(ctx, node, loop)
}
