package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type whereFunction struct {
	interpreter.BaseFunction
}

func newWhereFunction() *whereFunction {
	return &whereFunction{
		interpreter.NewBaseFunction("where", 0, 1, 1),
	}
}

func (f *whereFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, loop interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var filtered interpreter.ColModifier
	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if res != nil {
			if b, ok := res.(interpreter.BooleanAccessor); !ok {
				return nil, fmt.Errorf("Filter expression must return boolean, but returned %T", res)
			} else if b.Bool() {
				if filtered == nil {
					filtered = ctx.NewCol()
				}
				filtered.Add(this)
			}
		}
	}

	return filtered, nil
}

type selectFunction struct {
	interpreter.BaseFunction
}

func newSelectFunction() *selectFunction {
	return &selectFunction{
		interpreter.NewBaseFunction("select", 0, 1, 1),
	}
}

func (f *selectFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, loop interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var projected interpreter.ColModifier
	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if res != nil {
			if projected == nil {
				projected = ctx.NewCol()
			}
			if c, ok := res.(interpreter.ColAccessor); ok {
				projected.AddAll(c)
			} else {
				projected.Add(res)
			}
		}
	}

	return projected, nil
}

type repeatFunction struct {
	interpreter.BaseFunction
}

var repeatFunc = &repeatFunction{
	interpreter.NewBaseFunction("repeat", 0, 1, 1),
}

func (f *repeatFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, loop interpreter.Looper) (interface{}, error) {
	projected := ctx.NewCol()
	err := repeat(ctx, node, loop, projected)
	if err != nil || projected.Empty() {
		projected = nil
	}

	return projected, err
}

func repeat(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper, projected interpreter.ColModifier) error {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil
	}

	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return err
		}
		if res != nil {
			err := repeatRecursively(ctx, res, loop, projected)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func repeatRecursively(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper, projected interpreter.ColModifier) error {
	if col, ok := node.(interpreter.ColAccessor); ok {
		count := col.Count()
		for i := 0; i < count; i++ {
			n := col.Get(i)
			if n != nil {
				added := projected.AddUnique(n)
				if added {
					err := repeat(ctx, n, interpreter.NewLoopWithIndex(loop.Evaluator(), i), projected)
					if err != nil {
						return err
					}
				}
			}
		}
	} else if projected.AddUnique(node) {
		err := repeat(ctx, node, interpreter.NewLoop(loop.Evaluator()), projected)
		if err != nil {
			return err
		}
	}

	return nil
}

type ofTypeFunction struct {
	interpreter.BaseFunction
}

func newOfTypeFunction() *ofTypeFunction {
	return &ofTypeFunction{
		interpreter.NewBaseFunction("ofType", -1, 1, 1),
	}
}

func (f *ofTypeFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	var typeSpec interpreter.StringAccessor
	var ok bool
	if typeSpec, ok = unwrapCollection(args[0]).(interpreter.StringAccessor); !ok {
		return nil, fmt.Errorf("Expected string argument to ofType, but got %T", args[0])
	}

	var typeName interpreter.FQTypeNameAccessor
	var err error
	if typeName, err = interpreter.ParseFQTypeName(typeSpec.String()); err != nil {
		return nil, fmt.Errorf("Invalid type specifier '%s'", typeSpec.String())
	}

	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var filtered interpreter.ColModifier
	adapter := ctx.ModelAdapter()

	for i := 0; i < count; i++ {
		n := col.Get(i)
		if interpreter.HasModelType(adapter, n, typeName) {
			if filtered == nil {
				filtered = ctx.NewCol()
			}
			filtered.Add(n)
		}
	}

	return filtered, nil
}
