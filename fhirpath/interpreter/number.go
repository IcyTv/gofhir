package interpreter

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

var decimalTen = decimal.NewFromInt32(10)

type ArithmeticOps byte

const (
	AdditionOp       ArithmeticOps = '+'
	SubtractionOp    ArithmeticOps = '-'
	MultiplicationOp ArithmeticOps = '*'
	DivisionOp       ArithmeticOps = '/'
	DivOp            ArithmeticOps = 'D'
	ModOp            ArithmeticOps = 'M'
)

type DecimalValueAccessor interface {
	AnyAccessor
	Value() DecimalAccessor
	WithValue(node NumberAccessor) DecimalValueAccessor
	ArithmeticOpSupported(op ArithmeticOps) bool
}

type ArithmeticApplier interface {
	Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error)
	Abs() DecimalValueAccessor
}

type NumberAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier
	Ceiling() NumberAccessor
	Exp() NumberAccessor
	Floor() NumberAccessor
	Ln() (NumberAccessor, error)
	Log(base NumberAccessor) (NumberAccessor, error)
	Power(exponent NumberAccessor) (NumberAccessor, bool)
	Round(precision int32) (NumberAccessor, error)
	Sqrt() (NumberAccessor, bool)
	Truncate(precision int32) NumberAccessor
	Int() int32
	Int64() int64
	Float32() float32
	Float64() float64
	BigFloat() *big.Float
	Decimal() decimal.Decimal
	One() bool
	Positive() bool
	HasFraction() bool
}

func leastPrecisionDecimal(d1 decimal.Decimal, d2 decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	p1, p2 := decimalPrecision(d1), decimalPrecision(d2)
	if p1 == p2 {
		return d1, d2
	}

	if p1 < p2 {
		return d1, d2.Truncate(p1)
	}
	return d1.Truncate(p2), d2
}

func decimalPrecision(d decimal.Decimal) int32 {
	precision := -d.Exponent()
	if precision <= 0 {
		return 0
	}

	v := d.Mul(decimal.NewFromInt(int64(math.Pow(10.0, float64(precision)))))
	for precision > 0 {
		m := v.Mod(decimalTen)
		if !m.IsZero() {
			return precision
		}
		v = v.Div(decimalTen)
		precision--
	}
	return precision
}
