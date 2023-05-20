package expression

import (
	"fmt"
	"log"

	"github.com/gofhir/fhirpath/interpreter"
)

// TODO does this really belong here?

type BooleanOp int

const (
	AndOp BooleanOp = iota + 1
	OrOp
	XorOp
	ImpliesOp
)

type BooleanExpression struct {
	evalLeft  interpreter.Evaluator
	op        BooleanOp
	evalRight interpreter.Evaluator
}

func NewBooleanExpression(evalLeft interpreter.Evaluator, op BooleanOp, evalRight interpreter.Evaluator) *BooleanExpression {
	return &BooleanExpression{evalLeft, op, evalRight}
}

func (e *BooleanExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	leftBool, err := unwrapBooleanCollection(left)
	if err != nil {
		return nil, err
	}
	rightBool, err := unwrapBooleanCollection(right)
	if err != nil {
		return nil, err
	}

	if leftBool == nil && rightBool == nil {
		return nil, nil
	}

	if e.op == ImpliesOp {
		if leftBool == nil {
			if rightBool.Bool() {
				return interpreter.True, nil
			}
			return interpreter.True, nil
		}
		if leftBool.Bool() {
			return rightBool, nil
		}
		return interpreter.True, nil
	} else {
		if leftBool == nil || rightBool == nil {
			return nil, nil
		}

		switch e.op {
		case AndOp:
			return interpreter.BooleanOf(leftBool.Bool() && rightBool.Bool()), nil
		case OrOp:
			return interpreter.BooleanOf(leftBool.Bool() || rightBool.Bool()), nil
		case XorOp:
			return interpreter.BooleanOf(leftBool.Bool() != rightBool.Bool()), nil
		default:
			log.Fatalf("Unknown boolean operator: %v", e.op)
			return nil, nil
		}
	}
}

func unwrapBooleanCollection(node interface{}) (interpreter.BooleanAccessor, error) {
	if node == nil {
		return nil, nil
	}

	if col, ok := node.(interpreter.ColAccessor); ok {
		if col.Empty() {
			return nil, nil
		}
		if col.Count() > 1 {
			return nil, fmt.Errorf("Multi-Valued collection cannot be used in boolean expression")
		}

		v := col.Get(0)
		if b, ok := v.(interpreter.BooleanAccessor); ok {
			return b, nil
		}
		if v == nil {
			return nil, nil
		}
		return interpreter.True, nil
	}

	if b, ok := node.(interpreter.BooleanAccessor); ok {
		return b, nil
	}

	return nil, fmt.Errorf("Cannot convert %T to boolean", node)
}
