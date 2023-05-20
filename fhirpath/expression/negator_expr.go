package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type NegatorExpression struct {
	evaluator interpreter.Evaluator
}

func NewNegatorExpression(evaluator interpreter.Evaluator) *NegatorExpression {
	return &NegatorExpression{evaluator}
}

func (e *NegatorExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	data, err := e.evaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	negator, ok := data.(interpreter.Negator)
	if !ok {
		return nil, fmt.Errorf("Cannot negate non-negatable type: %T", data)
	}

	return negator.Negate(), nil
}
