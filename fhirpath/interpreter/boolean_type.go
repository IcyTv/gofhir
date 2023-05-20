package interpreter

import "fmt"

var BooleanTypeSpec = newAnyTypeSpec("Boolean")

var booleanTypeInfo = NewSimpleTypeInfo(namespaceNameString, NewString("Boolean"), anyTypeNameString)

type booleanType struct {
	baseAnyType
	value bool
}

type BooleanAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	Negator
	Bool() bool
}

var (
	True  = NewBooleanWithSource(true, nil)
	False = NewBooleanWithSource(false, nil)
)

func BooleanOf(value bool) BooleanAccessor {
	if value {
		return True
	} else {
		return False
	}
}

func NewBoolean(value bool) BooleanAccessor {
	return NewBooleanWithSource(value, nil)
}

func NewBooleanWithSource(value bool, source interface{}) BooleanAccessor {
	return newBoolean(value, source)
}

func newBoolean(value bool, source interface{}) BooleanAccessor {
	return &booleanType{
		baseAnyType{source},
		value,
	}
}

func ParseBoolean(value string) (BooleanAccessor, error) {
	switch value {
	case "true":
		return True, nil
	case "false":
		return False, nil
	}
	return nil, fmt.Errorf("Not a boolean value: %s", value)
}

func (t *booleanType) DataType() DataTypes {
	return BooleanDataType
}

func (t *booleanType) Bool() bool {
	return t.value
}

func (t *booleanType) String() string {
	if t.value {
		return "true"
	} else {
		return "false"
	}
}

func (e *booleanType) TypeSpec() TypeSpecAccessor {
	return BooleanTypeSpec
}

func (e *booleanType) TypeInfo() TypeInfoAccessor {
	return booleanTypeInfo
}

func (t *booleanType) Equal(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}
	return t.Bool() == node.(BooleanAccessor).Bool()
}

func (t *booleanType) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *booleanType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	} else {
		if t.value == comparator.(BooleanAccessor).Bool() {
			return 0, Evaluated
		}
		if !t.value {
			return -1, Evaluated
		}
		return 1, Evaluated
	}
}

func (t *booleanType) Negate() AnyAccessor {
	return BooleanOf(!t.value)
}
