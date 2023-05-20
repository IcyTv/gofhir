package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type absFunction struct {
	interpreter.BaseFunction
}

func newAbsFunction() *absFunction {
	return &absFunction{
		interpreter.NewBaseFunction("abs", -1, 0, 0),
	}
}

func (f *absFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := arithmeticNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Abs(), nil
}

type ceilingFunction struct {
	interpreter.BaseFunction
}

func newCeilingFunction() *ceilingFunction {
	return &ceilingFunction{
		interpreter.NewBaseFunction("ceiling", -1, 0, 0),
	}
}

func (f *ceilingFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := numberNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Ceiling(), nil
}

type expFunction struct {
	interpreter.BaseFunction
}

func newExpFunction() *expFunction {
	return &expFunction{
		interpreter.NewBaseFunction("exp", -1, 0, 0),
	}
}

func (f *expFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := numberNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Exp(), nil
}

type floorFunction struct {
	interpreter.BaseFunction
}

func newFloorFunction() *floorFunction {
	return &floorFunction{
		interpreter.NewBaseFunction("floor", -1, 0, 0),
	}
}

func (f *floorFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := numberNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Floor(), nil
}

type lnFunction struct {
	interpreter.BaseFunction
}

func newLnFunction() *lnFunction {
	return &lnFunction{
		interpreter.NewBaseFunction("ln", -1, 0, 0),
	}
}

func (f *lnFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := numberNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Ln()
}

type logFunction struct {
	interpreter.BaseFunction
}

func newLogFunction() *logFunction {
	return &logFunction{
		interpreter.NewBaseFunction("log", -1, 1, 1),
	}
}

func (f *logFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	a, err := numberNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	b, err := numberNode(args[0])
	if b == nil || err != nil {
		return nil, err
	}

	return a.Log(b)
}

type powerFunction struct {
	interpreter.BaseFunction
}

func newPowerFunction() *powerFunction {
	return &powerFunction{
		interpreter.NewBaseFunction("power", -1, 1, 1),
	}
}

func (f *powerFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	p, err := numberNode(args[0])
	if p == nil || err != nil {
		return nil, err
	}

	r, ok := n.Power(p)
	if !ok {
		return nil, nil
	}
	return r, nil
}

type roundFunction struct {
	interpreter.BaseFunction
}

func newRoundFunction() *roundFunction {
	return &roundFunction{
		interpreter.NewBaseFunction("round", -1, 0, 1),
	}
}

func (f *roundFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	precision := int32(0)
	if len(args) > 0 {
		p, err := integerNode(args[0])
		if p == nil || err != nil {
			return nil, err
		}
		precision = p.Int()
	}

	r, err := n.Round(precision)
	if err != nil {
		return nil, err
	}

	if i, ok := r.(interpreter.IntegerAccessor); ok {
		return interpreter.NewDecimal(i.Decimal()), nil
	}
	return r, nil
}

type sqrtFunction struct {
	interpreter.BaseFunction
}

func newSqrtFunction() *sqrtFunction {
	return &sqrtFunction{
		interpreter.NewBaseFunction("sqrt", -1, 0, 0),
	}
}

func (f *sqrtFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	r, ok := n.Sqrt()
	if !ok {
		return nil, nil
	}
	return r, nil
}

type truncateFunction struct {
	interpreter.BaseFunction
}

func newTruncateFunction() *truncateFunction {
	return &truncateFunction{
		BaseFunction: interpreter.NewBaseFunction("truncate", -1, 0, 0),
	}
}

func (f *truncateFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	t := n.Truncate(0)
	if i, ok := t.(interpreter.IntegerAccessor); ok {
		return i, nil
	}
	return interpreter.NewInteger(t.Int()), nil
}

func arithmeticNode(node interface{}) (interpreter.ArithmeticApplier, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(interpreter.ArithmeticApplier); !ok {
		return nil, fmt.Errorf("arithmetic cannot be applied: %T", value)
	} else {
		return a, nil
	}
}

func numberNode(node interface{}) (interpreter.NumberAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(interpreter.NumberAccessor); !ok {
		return nil, fmt.Errorf("not a number: %T", value)
	} else {
		return a, nil
	}
}

func integerNode(node interface{}) (interpreter.IntegerAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(interpreter.IntegerAccessor); !ok {
		return nil, fmt.Errorf("not an integer: %T", value)
	} else {
		return a, nil
	}
}
