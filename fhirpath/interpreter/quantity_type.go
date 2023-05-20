package interpreter

import (
	"fmt"
	"strings"
)

var UCUMSystemURI = NewString("http://unitsofmeasure.org")
var QuantityTypeSpec = newAnyTypeSpec("Quantity")

var quantityTypeInfo = NewClassInfo(namespaceNameString, NewString("Quantity"), NewString("System.Any"),
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("value"), NewString(DecimalTypeSpec.String()), nil),
		NewClassInfoElement(NewString("unit"), NewString(StringTypeSpec.String()), nil),
	}),
)

type quantityType struct {
	baseAnyType
	value DecimalAccessor
	unit  StringAccessor
}

type QuantityAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier

	Unit() StringAccessor
	ToUnit(unit StringAccessor) QuantityAccessor
}

func NewQuantity(value DecimalAccessor, unit StringAccessor) QuantityAccessor {
	return NewQuantityWithSource(value, unit, nil)
}

func NewQuantityWithSource(value DecimalAccessor, unit StringAccessor, source interface{}) QuantityAccessor {
	if value == nil {
		panic("value cannot be nil")
	}
	return &quantityType{
		baseAnyType{source},
		value,
		unit,
	}
}

func (t *quantityType) DataType() DataTypes {
	return QuantityDataType
}

func (t *quantityType) Value() DecimalAccessor {
	return t.value
}

func (t *quantityType) Unit() StringAccessor {
	return t.unit
}

func (t *quantityType) WithValue(node NumberAccessor) DecimalValueAccessor {
	var value DecimalAccessor
	if node == nil {
		return nil
	} else if node.DataType() == DecimalDataType {
		value = node.(DecimalAccessor)
	} else {
		value = NewDecimal(node.Decimal())
	}
	return NewQuantity(value, t.Unit())
}

func (t *quantityType) ArithmeticOpSupported(op ArithmeticOps) bool {
	return op == AdditionOp || op == SubtractionOp || op == MultiplicationOp || op == DivisionOp
}

func (t *quantityType) Negate() AnyAccessor {
	return NewQuantity(t.Value().Negate().(DecimalAccessor), t.Unit())
}

func (e *quantityType) TypeSpec() TypeSpecAccessor {
	return QuantityTypeSpec
}

func (e *quantityType) TypeInfo() TypeInfoAccessor {
	return quantityTypeInfo
}

func (t *quantityType) Equal(node interface{}) bool {
	return quantityValueEqual(t, node, false)
}

func (t *quantityType) Equivalent(node interface{}) bool {
	return quantityValueEqual(t, node, true)
}

func quantityValueEqual(t QuantityAccessor, node interface{}, equivalent bool) bool {
	if q, ok := node.(QuantityAccessor); ok {
		if Equal(t.Unit(), q.Unit()) {
			return t.Value().Equal(q.Value())
		}

		u1, exp1 := QuantityUnitWithNameString(t.Unit())
		u2, exp2 := QuantityUnitWithNameString(q.Unit())
		if exp1 != exp2 {
			return false
		}
		v1, v2, u := ConvertUnitToBase(t.Value(), u1, exp1, q.Value(), u2, exp2, !equivalent)
		if u != nil {
			return Equal(v1, v2)
		}
	} else if d, ok := node.(DecimalValueAccessor); ok {
		v1 := t.Value()
		v2 := d.Value()
		if equivalent {
			return Equivalent(v1, v2)
		}
		return Equal(v1, v2)
	}
	return false
}

func (t *quantityType) Compare(comparator Comparator) (int, OperatorStatus) {
	if q, ok := comparator.(QuantityAccessor); ok {
		if !Equal(t.Unit(), q.Unit()) {
			u1, exp1 := QuantityUnitWithNameString(t.Unit())
			u2, exp2 := QuantityUnitWithNameString(q.Unit())
			if exp1 == exp2 {
				v1, v2, u := ConvertUnitToBase(t.Value(), u1, exp1, q.Value(), u2, exp2, false)
				if u != nil {
					return decimalValueCompare(v1, v2)
				}
			}

			return -1, Empty
		}
	}

	return decimalValueCompare(t.value, comparator)
}

func (t *quantityType) ToUnit(unit StringAccessor) QuantityAccessor {
	u2, exp2 := QuantityUnitWithNameString(unit)
	if u2 == nil {
		return nil
	}

	u1, exp1 := QuantityUnitWithNameString(t.Unit())
	if u1 == nil || exp1 != exp2 {
		return nil
	}

	if u1.Equal(u2) {
		return t
	}

	u := u1.CommonBase(u2, true)
	if u == nil {
		return nil
	}

	f1, f2 := u1.Factor(u, exp1), u2.Factor(u, exp2)
	v, _ := t.Value().Calc(f1, MultiplicationOp)
	v, _ = v.Value().Calc(f2, DivisionOp)

	val := v.Value()
	return NewQuantity(val, u2.NameWithExp(val, exp2))
}

var builtinUnits = map[string]struct{}{
	"year":         {},
	"years":        {},
	"month":        {},
	"months":       {},
	"week":         {},
	"weeks":        {},
	"day":          {},
	"days":         {},
	"hour":         {},
	"hours":        {},
	"minute":       {},
	"minutes":      {},
	"second":       {},
	"seconds":      {},
	"millisecond":  {},
	"milliseconds": {},
}

func (t *quantityType) String() string {
	var b strings.Builder
	b.Grow(32)
	b.WriteString(t.value.String())
	if t.unit != nil {
		b.WriteByte(' ')
		if _, ok := builtinUnits[t.unit.String()]; !ok {
			b.WriteByte('\'')
			b.WriteString(t.unit.String())
			b.WriteByte('\'')
		} else {
			b.WriteString(t.unit.String())
		}
	}
	return b.String()
}

func (t *quantityType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if t.value == nil || operand == nil {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("Arithmetic Operator not supported %c", op)
	}

	var valLeft, valRight DecimalAccessor
	var unit QuantityUnitAccessor
	var exp int
	if q, ok := operand.(QuantityAccessor); !ok {
		valLeft, valRight = t.Value(), operand.Value()
		unit, exp = QuantityUnitWithNameString(t.Unit())
	} else {
		var err error
		valLeft, valRight, unit, exp, err = mergeQuantityUnits(t, q, op)
		if err != nil {
			return nil, err
		}
	}

	value, _ := valLeft.Calc(valRight, op)
	return NewQuantity(value.Value(), unit.NameWithExp(value.Value(), exp)), nil
}

func (t *quantityType) Abs() DecimalValueAccessor {
	return NewQuantity(t.Value().Abs().(DecimalAccessor), t.Unit())
}

func mergeQuantityUnits(l QuantityAccessor, r QuantityAccessor, op ArithmeticOps) (DecimalAccessor, DecimalAccessor, QuantityUnitAccessor, int, error) {
	leftVal, rightVal := l.Value(), r.Value()
	leftUnit, leftExp := QuantityUnitWithNameString(l.Unit())
	rightUnit, rightExp := QuantityUnitWithNameString(l.Unit())

	var unit QuantityUnitAccessor
	if leftUnit == nil && rightUnit == nil {
		return leftVal, rightVal, EmptyQuantityUnit, 1, nil
	}
	if leftUnit != nil && leftUnit.Equal(rightUnit) {
		unit = leftUnit
	} else {
		leftVal, rightVal, unit = ConvertUnitToMostGranular(
			leftVal, leftUnit, leftExp, rightVal, rightUnit, rightExp, true,
		)
		if unit == nil {
			return nil, nil, nil, 1, fmt.Errorf("Units are not equal: %s != %s", leftUnit, rightUnit)
		}
	}

	exp := leftExp
	switch op {
	case AdditionOp, SubtractionOp:
		if leftExp != rightExp {
			return nil, nil, nil, 1, fmt.Errorf("Unit exponents are unequal: %d != %d", leftExp, rightExp)
		}
	case MultiplicationOp:
		exp = leftExp + rightExp
	case DivisionOp:
		exp = leftExp - rightExp
	}

	if exp < 1 || exp > 3 {
		return nil, nil, nil, 1, fmt.Errorf("Resulting unit exponent must be between 1 and 3, is %d", exp)
	}

	return leftVal, rightVal, unit, exp, nil
}
