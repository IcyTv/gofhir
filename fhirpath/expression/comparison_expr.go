package expression

import (
	"fmt"
	"log"

	"github.com/gofhir/fhirpath/interpreter"
)

type ComparisonOp int

const (
	LessThanOp ComparisonOp = iota + 1
	LessOrEqualThanOp
	GreaterThanOp
	GreaterOrEqualThanOp
)

type ComparisonExpression struct {
	evalLeft  interpreter.Evaluator
	op        ComparisonOp
	evalRight interpreter.Evaluator
}

func NewComparisonExpression(evalLeft interpreter.Evaluator, op ComparisonOp, evalRight interpreter.Evaluator) *ComparisonExpression {
	return &ComparisonExpression{evalLeft, op, evalRight}
}

func (e *ComparisonExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
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
		return nil, nil
	}

	var ok bool
	var leftCmp, rightCmp interpreter.Comparator
	if leftCmp, ok = left.(interpreter.Comparator); !ok {
		return nil, fmt.Errorf("Operand cannot be used for comparison: %T", left)
	}
	if rightCmp, ok = right.(interpreter.Comparator); !ok {
		return nil, fmt.Errorf("Operand cannot be used for comparison: %T", right)
	}

	res, status := leftCmp.Compare(rightCmp)
	if status == interpreter.Empty {
		return nil, nil
	}
	if status != interpreter.Evaluated {
		return nil, fmt.Errorf("Operands cannot be compared: %T <> %T", left, right)
	}

	var b bool
	switch e.op {
	case LessOrEqualThanOp:
		b = res <= 0
	case LessThanOp:
		b = res < 0
	case GreaterOrEqualThanOp:
		b = res >= 0
	case GreaterThanOp:
		b = res > 0
	default:
		log.Fatalf("Unknown comparison operator: %v", e.op)
	}

	return interpreter.BooleanOf(b), nil
}
