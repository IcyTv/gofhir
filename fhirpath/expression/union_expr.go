package expression

import "github.com/gofhir/fhirpath/interpreter"

type UnionExpression struct {
	evalLeft  interpreter.Evaluator
	evalRight interpreter.Evaluator
}

func NewUnionExpression(evalLeft, evalRight interpreter.Evaluator) *UnionExpression {
	return &UnionExpression{evalLeft, evalRight}
}

func (e *UnionExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	return uniteCollections(ctx, left, right), nil
}
