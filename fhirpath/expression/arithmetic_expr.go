package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type ArithmeticExpression struct {
	evalLeft  interpreter.Evaluator
	op        interpreter.ArithmeticOps
	evalRight interpreter.Evaluator
}

func NewArithmeticExpression(evalLeft interpreter.Evaluator, op interpreter.ArithmeticOps, evalRight interpreter.Evaluator) *ArithmeticExpression {
	return &ArithmeticExpression{evalLeft, op, evalRight}
}

func (e *ArithmeticExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
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

	leftOperand, ok := left.(interpreter.ArithmeticApplier)
	if !ok {
		return applyNonNumericArithmetic(left, e.op, right)
	}
	rightOperand, ok := right.(interpreter.DecimalValueAccessor)
	if !ok {
		return applyNonNumericArithmetic(left, e.op, right)
	}

	return leftOperand.Calc(rightOperand, e.op)
}

func applyNonNumericArithmetic(left interface{}, op interpreter.ArithmeticOps, right interface{}) (interpreter.AnyAccessor, error) {
	if op == interpreter.AdditionOp || op == interpreter.SubtractionOp {
		t, err := applyTemporalArithmetic(left, right, op == interpreter.SubtractionOp)
		if err != nil {
			return nil, err
		}
		if t != nil {
			return t, nil
		}
	}

	if op == interpreter.AdditionOp {
		if s := applyStringArithmetic(left, right); s != nil {
			return s, nil
		}
	}

	return nil, fmt.Errorf("Operands %T and %T do not support arithmetic operation %c", left, right, op)
}

func applyStringArithmetic(left, right interface{}) interpreter.StringAccessor {
	var ok bool
	var leftString, rightString interpreter.Stringifier
	if leftString, ok = left.(interpreter.Stringifier); !ok {
		return nil
	}
	if rightString, ok = right.(interpreter.Stringifier); !ok {
		return nil
	}

	return interpreter.NewString(leftString.String() + rightString.String())
}

func applyTemporalArithmetic(left, right interface{}, negate bool) (interpreter.TemporalAccessor, error) {
	var ok bool
	var temporal interpreter.TemporalAccessor
	var quantity interpreter.QuantityAccessor
	if temporal, ok = left.(interpreter.TemporalAccessor); !ok {
		return nil, nil
	}
	if quantity, ok = right.(interpreter.QuantityAccessor); !ok {
		return nil, fmt.Errorf("Only a quantity may be added to a temporal value %T", right)
	}

	if negate {
		quantity = quantity.Negate().(interpreter.QuantityAccessor)
	}

	return temporal.Add(quantity)
}
