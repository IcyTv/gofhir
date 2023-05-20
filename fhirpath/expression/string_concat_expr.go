package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type StringConcatExpression struct {
	evalLeft  interpreter.Evaluator
	evalRight interpreter.Evaluator
}

func NewStringConcatExpression(evalLeft, evalRight interpreter.Evaluator) *StringConcatExpression {
	return &StringConcatExpression{evalLeft, evalRight}
}

func (e *StringConcatExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	left, right = unwrapCollection(left), unwrapCollection(right)
	if left == nil || right == nil {
		return interpreter.EmptyString, nil
	}

	var ok bool
	var leftStr, rightStr interpreter.Stringifier
	if left != nil {
		if leftStr, ok = left.(interpreter.Stringifier); !ok {
			return nil, fmt.Errorf("Left operand is not string: %T", left)
		}
	}
	if right != nil {
		if rightStr, ok = right.(interpreter.Stringifier); !ok {
			return nil, fmt.Errorf("Right operand is not string: %T", right)
		}
	}

	if leftStr == nil {
		return rightStr, nil
	}
	if rightStr == nil {
		return leftStr, nil
	}

	return interpreter.NewString(leftStr.String() + rightStr.String()), nil
}
