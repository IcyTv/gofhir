package expression

import "github.com/gofhir/fhirpath/interpreter"

type unionFunction struct {
	interpreter.BaseFunction
}

func newUnionFunction() *unionFunction {
	return &unionFunction{
		interpreter.NewBaseFunction("union", -1, 1, 1),
	}
}

func (f *unionFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	return uniteCollections(ctx, node, args[0]), nil
}

type combineFunction struct {
	interpreter.BaseFunction
}

func newCombineFunction() *combineFunction {
	return &combineFunction{
		interpreter.NewBaseFunction("combine", -1, 1, 1),
	}
}

func (f *combineFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	return combineCollections(ctx, node, args[0]), nil
}
