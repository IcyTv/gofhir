package expression

import "github.com/gofhir/fhirpath/interpreter"

type EqualityExpression struct {
	not                 bool
	equivalent          bool
	evalLeft, evalRight interpreter.Evaluator
}

func NewEqualityExpression(not bool, equivalent bool, evalLeft, evalRight interpreter.Evaluator) interpreter.Evaluator {
	return &EqualityExpression{not, equivalent, evalLeft, evalRight}
}

func (e *EqualityExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	res, err := e.evaluateInternally(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	if e.not {
		return res.(interpreter.BooleanAccessor).Negate(), nil
	}

	return res, nil
}

func (e *EqualityExpression) evaluateInternally(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
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
		if e.equivalent {
			return interpreter.BooleanOf(interpreter.ModelEquivalent(ctx.ModelAdapter(), left, right)), nil
		} else {
			return nil, nil
		}
	}

	r, match := e.stringsEqual(left, right)
	if match {
		return interpreter.BooleanOf(r), nil
	}

	if e.equivalent {
		r = interpreter.ModelEquivalent(ctx.ModelAdapter(), left, right)
	} else {
		r = interpreter.ModelEqual(ctx.ModelAdapter(), left, right)
		if !r && temporalPrecisionNotEqual(left, right) {
			return nil, nil
		}
	}

	return interpreter.BooleanOf(r), nil
}

func (e *EqualityExpression) stringsEqual(n1 interface{}, n2 interface{}) (equal bool, match bool) {
	var ok bool
	var s1, s2 interpreter.Stringifier
	if s1, ok = n1.(interpreter.Stringifier); !ok {
		return
	}
	if s2, ok = n2.(interpreter.Stringifier); !ok {
		return
	}

	if s1.DataType() != interpreter.StringDataType && s2.DataType() != interpreter.StringDataType {
		return
	}

	match = true
	if e.equivalent {
		equal = interpreter.NormalizedStringEqual(s1.String(), s2.String())
	} else {
		equal = s1.String() == s2.String()
	}
	return
}

func temporalPrecisionNotEqual(n1 interface{}, n2 interface{}) bool {
	var ok bool
	var t1, t2 interpreter.TemporalAccessor
	if t1, ok = n1.(interpreter.TemporalAccessor); !ok {
		return false
	}
	if t2, ok = n2.(interpreter.TemporalAccessor); !ok {
		return false
	}
	return !interpreter.TemporalPrecisionEqual(t1, t2)
}
