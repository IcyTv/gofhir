package expression

import "github.com/gofhir/fhirpath/interpreter"

type childrenFunction struct {
	interpreter.BaseFunction
}

var childrenFunc = &childrenFunction{
	interpreter.NewBaseFunction("children", -1, 0, 0),
}
var childrenFuncInvocation = newFunctionInvocation(childrenFunc, []interpreter.Evaluator{})

func (f *childrenFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	if node == nil {
		return nil, nil
	}

	if col, ok := node.(interpreter.ColAccessor); ok {
		count := col.Count()
		adapter := ctx.ModelAdapter()
		var children interpreter.ColModifier

		for i := 0; i < count; i++ {
			c := col.Get(i)
			if c != nil {
				ccol, err := adapter.Children(c)
				if err != nil {
					return nil, err
				}
				if ccol != nil && !ccol.Empty() {
					if children == nil {
						children = interpreter.NewCol(adapter)
					}
					children.AddAll(ccol)
				}
			}
		}

		return children, nil
	}

	return ctx.ModelAdapter().Children(node)
}

type descendantsFunction struct {
	interpreter.BaseFunction
}

func newDescendantsFunction() interpreter.FunctionExecutor {
	return &descendantsFunction{
		interpreter.NewBaseFunction("descendants", -1, 0, 0),
	}
}

func (f *descendantsFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	return repeatFunc.Execute(ctx, node, args, interpreter.NewLoop(childrenFuncInvocation))
}
