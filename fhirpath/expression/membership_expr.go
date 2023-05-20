package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type ContainsExpression struct {
	evalLeft  interpreter.Evaluator
	evalRight interpreter.Evaluator
	inverse   bool
}

func NewContainsExpression(evalLeft interpreter.Evaluator, evalRight interpreter.Evaluator, inverse bool) *ContainsExpression {
	return &ContainsExpression{evalLeft, evalRight, inverse}
}

func (e *ContainsExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if e.inverse {
		tmp := left
		left = right
		right = tmp
	}

	col := wrapCollection(ctx, left)
	val := unwrapCollection(right)
	if col == nil || col.Empty() {
		return interpreter.False, nil
	}
	if val == nil {
		return nil, nil
	}
	if interpreter.IsCol(val) {
		return nil, fmt.Errorf("Collection membership cannot be checked with value %T", val)
	}

	return interpreter.BooleanOf(col.Contains(val)), nil
}
