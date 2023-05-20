package interpreter

import "strings"

var StringTypeSpec = newAnyTypeSpec("String")
var stringTypeInfo = NewSimpleTypeInfo(namespaceNameString, NewString("String"), anyTypeNameString)
var EmptyString = newString("", nil)

type stringType struct {
	baseAnyType
	value string
}

type StringAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	Length() int
}

func StringOfNil(value string) StringAccessor {
	if len(value) == 0 {
		return nil
	}
	return NewString(value)
}

func StringOf(value string) StringAccessor {
	if len(value) == 0 {
		return EmptyString
	}
	return NewString(value)
}

func NewString(value string) StringAccessor {
	return NewStringWithSource(value, nil)
}

func NewStringWithSource(value string, source interface{}) StringAccessor {
	return newString(value, source)
}

func newString(value string, source interface{}) StringAccessor {
	return &stringType{
		baseAnyType{source},
		value,
	}
}

func (t *stringType) DataType() DataTypes {
	return StringDataType
}

func (t *stringType) String() string {
	return t.value
}

func (t *stringType) Length() int {
	return len(t.value)
}

func (e *stringType) TypeSpec() TypeSpecAccessor {
	return StringTypeSpec
}

func (e *stringType) TypeInfo() TypeInfoAccessor {
	return stringTypeInfo
}

func (t *stringType) Equal(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}

	return t.String() == node.(Stringifier).String()
}

func (t *stringType) Equivalent(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}

	return NormalizedStringEqual(t.String(), node.(Stringifier).String())
}

func (t *stringType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	} else {
		return strings.Compare(t.value, comparator.(StringAccessor).String()), Evaluated
	}
}
