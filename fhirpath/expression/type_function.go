package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type asFunction struct {
	interpreter.BaseFunction
}

func newAsFunction() *asFunction {
	return &asFunction{
		interpreter.NewBaseFunction("as", -1, 1, 1),
	}
}

func (f *asFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return nil, err
	}

	fqName, err := interpreter.ParseFQTypeName(name.String())
	if err != nil {
		return nil, err
	}

	item := unwrapCollection(node)
	if _, ok := item.(interpreter.ColAccessor); ok {
		return nil, fmt.Errorf("as operator cannot be applied on a collection")
	}

	return interpreter.CastModelType(ctx.ModelAdapter(), item, fqName)
}

type isFunction struct {
	interpreter.BaseFunction
}

func newIsFunction() *isFunction {
	return &isFunction{
		interpreter.NewBaseFunction("is", -1, 1, 1),
	}
}

func (f *isFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return nil, err
	}

	fqName, err := interpreter.ParseFQTypeName(name.String())
	if err != nil {
		return nil, err
	}

	item := unwrapCollection(node)
	if _, ok := item.(interpreter.ColAccessor); ok {
		return nil, fmt.Errorf("is operator cannot be applied on a collection")
	}

	return interpreter.BooleanOf(interpreter.HasModelType(ctx.ModelAdapter(), item, fqName)), nil
}
