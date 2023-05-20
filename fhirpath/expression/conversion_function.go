package expression

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofhir/fhirpath/interpreter"
)

type iifFunction struct {
	interpreter.BaseFunction
}

func newIIfFunction() *iifFunction {
	return &iifFunction{
		interpreter.NewBaseFunction("iif", -1, 2, 3),
	}
}

func (f *iifFunction) Execute(_ interpreter.ContextAccessor, _ interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	criterion := unwrapCollection(args[0])
	var criterionValue bool
	if criterion != nil {
		if b, ok := criterion.(interpreter.BooleanAccessor); !ok {
			return nil, fmt.Errorf("Criterion cannot be used as boolean: %T", criterion)
		} else {
			criterionValue = b.Bool()
		}
	}

	var res interface{}
	if criterionValue {
		res = args[1]
	} else if len(args) > 2 {
		res = args[2]
	}

	return res, nil
}

type toBooleanFunction struct {
	interpreter.BaseFunction
}

var toBooleanFunc = &toBooleanFunction{
	interpreter.NewBaseFunction("toBoolean", -1, 0, 0),
}

func (f *toBooleanFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if b, ok := any.(interpreter.BooleanAccessor); ok {
		return b, nil
	}
	if n, ok := any.(interpreter.NumberAccessor); ok {
		f := n.Float64()
		if f == 1.0 {
			return interpreter.True, nil
		}
		if f == 0.0 {
			return interpreter.False, nil
		}
	} else if s, ok := any.(interpreter.StringAccessor); ok {
		s := strings.ToLower(s.String())
		switch s {
		case "true", "t", "yes", "y", "1", "1.0":
			return interpreter.True, nil
		case "false", "f", "no", "n", "0", "0.0":
			return interpreter.False, nil
		}
	}

	// TODO how is this not an error?

	return nil, nil
}

type convertsToBooleanFunction struct {
	convertsToFunction
}

func newConvertsToBooleanFunction() *convertsToBooleanFunction {
	return &convertsToBooleanFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToBoolean", -1, 0, 0),
			toBooleanFunc,
		},
	}
}

type toIntegerFunction struct {
	interpreter.BaseFunction
}

var toIntegerFunc = &toIntegerFunction{
	interpreter.NewBaseFunction("toInteger", -1, 0, 0),
}

func (f *toIntegerFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	switch any.DataType() {
	case interpreter.IntegerDataType:
		return any, nil
	case interpreter.BooleanDataType:
		if any.(interpreter.BooleanAccessor).Bool() {
			return interpreter.NewInteger(1), nil
		}
		return interpreter.NewInteger(0), nil
	case interpreter.StringDataType:
		i, err := interpreter.ParseInteger(any.(interpreter.StringAccessor).String())
		if err == nil {
			return i, nil
		}
	}

	return nil, nil
}

type convertsToIntegerFunction struct {
	convertsToFunction
}

func newConvertsToIntegerFunction() *convertsToIntegerFunction {
	return &convertsToIntegerFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToInteger", -1, 0, 0),
			toIntegerFunc,
		},
	}
}

type toDecimalFunction struct {
	interpreter.BaseFunction
}

var toDecimalFunc = &toDecimalFunction{
	interpreter.NewBaseFunction("toDecimal", -1, 0, 0),
}

func (f *toDecimalFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if n, ok := any.(interpreter.NumberAccessor); ok {
		return n.Value(), nil
	}

	switch any.DataType() {
	case interpreter.BooleanDataType:
		if any.(interpreter.BooleanAccessor).Bool() {
			return interpreter.NewDecimalInt(1), nil
		} else {
			return interpreter.NewDecimalInt(0), nil
		}
	case interpreter.StringDataType:
		d, err := interpreter.ParseDecimal(any.(interpreter.StringAccessor).String())
		if err == nil {
			return d, nil
		}
	}

	return nil, nil
}

type convertsToDecimalFunction struct {
	convertsToFunction
}

func newConvertsToDecimalFunction() *convertsToDecimalFunction {
	return &convertsToDecimalFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToDecimal", -1, 0, 0),
			toDecimalFunc,
		},
	}
}

type toDateFunction struct {
	interpreter.BaseFunction
}

var toDateFunc = &toDateFunction{
	interpreter.NewBaseFunction("toDate", -1, 0, 0),
}

func (f *toDateFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d interpreter.DateAccessor
	if t, ok := any.(interpreter.DateTemporalAccessor); ok {
		d = t.Date()
	} else if s, ok := any.(interpreter.StringAccessor); ok {
		var err error
		d, err = interpreter.ParseDate(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToDateFunction struct {
	convertsToFunction
}

func newConvertsToDateFunction() *convertsToDateFunction {
	return &convertsToDateFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToDate", -1, 0, 0),
			toDateFunc,
		},
	}
}

type toDateTimeFunction struct {
	interpreter.BaseFunction
}

var toDateTimeFunc = &toDateTimeFunction{
	BaseFunction: interpreter.NewBaseFunction("toDateTime", -1, 0, 0),
}

func (f *toDateTimeFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d interpreter.DateTimeAccessor
	if t, ok := any.(interpreter.DateTemporalAccessor); ok {
		d = t.DateTime()
	} else if s, ok := any.(interpreter.StringAccessor); ok {
		var err error
		d, err = interpreter.ParseDateTime(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToDateTimeFunction struct {
	convertsToFunction
}

func newConvertsToDateTimeFunction() *convertsToDateTimeFunction {
	return &convertsToDateTimeFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: interpreter.NewBaseFunction("convertsToDateTime", -1, 0, 0),
			converter:    toDateTimeFunc,
		},
	}
}

// ==========================
// Quantity
// ==========================

type toQuantityFunction struct {
	interpreter.BaseFunction
}

var toQuantityFunc = &toQuantityFunction{
	interpreter.NewBaseFunction("toQuantity", -1, 0, 0),
}

var quantityStringPattern = regexp.MustCompile("^([+\\-]?\\d+(?:\\.\\d+)?)\\s*('(?:[^']+)'|(?:[a-zA-Z]+))?$")

func (f *toQuantityFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var q interpreter.QuantityAccessor
	if o, ok := any.(interpreter.QuantityAccessor); ok {
		q = o
	} else if s, ok := any.(interpreter.NumberAccessor); ok {
		v := s.Value()
		q = interpreter.NewQuantity(v, interpreter.DefaultQuantityUnit.Name(v))
	} else if s, ok := any.(interpreter.BooleanAccessor); ok {
		var v interpreter.DecimalAccessor
		if s.Bool() {
			v = interpreter.DecimalOne
		} else {
			v = interpreter.DecimalZero
		}
		q = interpreter.NewQuantity(v, interpreter.DefaultQuantityUnit.Name(v))
	} else if s, ok := any.(interpreter.StringAccessor); ok {
		m := quantityStringPattern.FindStringSubmatch(s.String())
		if m == nil {
			return nil, nil
		}
		v, _ := interpreter.ParseDecimal(m[1])
		u, exp := interpreter.QuantityUnitWithName(parseStringLiteral(m[2], stringDelimiterChar))
		if u == nil {
			q = interpreter.NewQuantity(v, nil)
		} else {
			q = interpreter.NewQuantity(v, u.NameWithExp(v, exp))
		}
	}

	if q != nil && len(args) > 0 {
		if s, ok := args[0].(interpreter.StringAccessor); !ok {
			return nil, fmt.Errorf("conversion unit is not string: %T", args[0])
		} else {
			q = q.ToUnit(s)
		}
	}

	return q, nil
}

type convertsToQuantityFunction struct {
	convertsToFunction
}

func newConvertsToQuantityFunction() *convertsToQuantityFunction {
	return &convertsToQuantityFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToQuantity", -1, 0, 0),
			toQuantityFunc,
		},
	}
}

// ==========================
// String
// ==========================

type toStringFunction struct {
	interpreter.BaseFunction
}

var toStringFunc = &toStringFunction{
	interpreter.NewBaseFunction("toString", -1, 0, 0),
}

func (f *toStringFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	return interpreter.NewString(any.(interpreter.Stringifier).String()), nil
}

type convertsToStringFunction struct {
	convertsToFunction
}

func newConvertsToStringFunction() *convertsToStringFunction {
	return &convertsToStringFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToString", -1, 0, 0),
			toStringFunc,
		},
	}
}

// ==========================
// Time
// ==========================

type toTimeFunction struct {
	interpreter.BaseFunction
}

var toTimeFunc = &toTimeFunction{
	interpreter.NewBaseFunction("toTime", -1, 0, 0),
}

func (f *toTimeFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d interpreter.TimeAccessor
	if t, ok := any.(interpreter.TimeAccessor); ok {
		d = t
	} else if s, ok := any.(interpreter.StringAccessor); ok {
		var err error
		d, err = interpreter.ParseTime(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToTimeFunction struct {
	convertsToFunction
}

func newConvertsToTimeFunction() *convertsToTimeFunction {
	return &convertsToTimeFunction{
		convertsToFunction{
			interpreter.NewBaseFunction("convertsToTime", -1, 0, 0),
			toTimeFunc,
		},
	}
}

// ==========================
// Utility functions/types
// ==========================

type convertsToFunction struct {
	interpreter.BaseFunction
	converter interpreter.FunctionExecutor
}

func (f *convertsToFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, loop interpreter.Looper) (interface{}, error) {
	if emptyCollection(node) {
		return interpreter.False, nil
	}

	res, err := f.converter.Execute(ctx, node, args, loop)
	if err != nil {
		return nil, err
	}

	if res != nil {
		return interpreter.True, nil
	}
	return interpreter.False, nil
}

func convertibleAny(node interface{}) (interpreter.AnyAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if any, ok := value.(interpreter.AnyAccessor); !ok {
		return nil, nil
	} else {
		if _, ok := any.(interpreter.ColAccessor); ok {
			return nil, fmt.Errorf("Collection with multiple items cannot be converted")
		}
		return any, nil
	}
}
