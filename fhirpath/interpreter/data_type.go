package interpreter

type DataTypes int

const UndefinedDataType DataTypes = 0x0001
const ColDataType DataTypes = 0x0002

const LiteralDataType DataTypes = 0x0200

const (
	BooleanDataType = iota + LiteralDataType
	IntegerDataType
	DecimalDataType
	StringDataType
	DateDataType
	DateTimeDataType
	TimeDataType
	QuantityDataType
)

var anyTypeNameString = NewString("System.Any")

type AnyAccessor interface {
	DataType() DataTypes
	TypeSpec() TypeSpecAccessor
	TypeInfo() TypeInfoAccessor
	Source() interface{}
	Equal(node interface{}) bool
	Equivalent(node interface{}) bool
}

type baseAnyType struct {
	source interface{}
}

func (a *baseAnyType) Source() interface{} {
	return a.source
}

func TypeEqual(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 != nil && n2 != nil && n1.DataType() == n2.DataType()
}

func Equal(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 == n2 || (n1 != nil && n2 != nil && n1.Equal(n2))
}

func Equivalent(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 == n2 || (n1 != nil && n2 != nil && n1.Equivalent(n2))
}

type OperatorStatus int

const (
	Inconvertible OperatorStatus = iota
	Empty
	Evaluated
)

type Comparator interface {
	AnyAccessor
	Compare(comparator Comparator) (int, OperatorStatus)
}

type Negator interface {
	AnyAccessor
	Negate() AnyAccessor
}

type Stringifier interface {
	AnyAccessor
	String() string
}
