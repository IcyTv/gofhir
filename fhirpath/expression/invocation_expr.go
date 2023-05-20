package expression

import "github.com/gofhir/fhirpath/interpreter"

type InvocationExpression struct {
	exprEvaluator       interpreter.Evaluator
	invocationEvaluator interpreter.Evaluator
}

func NewInvocationExpression(exprEvaluator, invocationEvaluator interpreter.Evaluator) *InvocationExpression {
	return &InvocationExpression{exprEvaluator, invocationEvaluator}
}

func (e *InvocationExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	exprNode, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	return e.invocationEvaluator.Evaluate(ctx, exprNode, loop)
}
