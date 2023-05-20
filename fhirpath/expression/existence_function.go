package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type emptyFunction struct {
	interpreter.BaseFunction
}

func newEmptyFunction() *emptyFunction {
	return &emptyFunction{
		interpreter.NewBaseFunction("empty", -1, 0, 0),
	}
}

func (f *emptyFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	if node == nil {
		return interpreter.True, nil
	}

	if col, ok := node.(interpreter.ColAccessor); ok {
		return interpreter.BooleanOf(col.Empty()), nil
	} else {
		return interpreter.False, nil
	}
}

type existsFunction struct {
	interpreter.BaseFunction
}

func newExistsFunction() *existsFunction {
	return &existsFunction{
		interpreter.NewBaseFunction("exists", 0, 0, 1),
	}
}

func (e *existsFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, loop interpreter.Looper) (interface{}, error) {
	if node == nil {
		return interpreter.False, nil
	}

	loopEvaluator := loop.Evaluator()
	col, ok := node.(interpreter.ColAccessor)
	if !ok {
		if loopEvaluator == nil {
			return interpreter.True, nil
		}
		col = ctx.NewColWithItem(node)
	}
	count := col.Count()

	found := false
	if loopEvaluator == nil {
		found = count > 0
	} else {
		for i := 0; i < count; i++ {
			for i := 0; i < count; i++ {
				this := col.Get(i)
				loop.IncIndex(this)

				res, err := loopEvaluator.Evaluate(ctx, this, loop)
				if err != nil {
					return nil, err
				}
				if res != nil {
					if b, ok := res.(interpreter.BooleanAccessor); !ok {
						return nil, fmt.Errorf("Expected boolean result from loop, got %T", res)
					} else if b.Bool() {
						return interpreter.True, nil
					}
				}
			}
		}
	}

	return interpreter.BooleanOf(found), nil
}

type allFunction struct {
	interpreter.BaseFunction
}

func newAllFunction() *allFunction {
	return &allFunction{
		interpreter.NewBaseFunction("all", 0, 1, 1),
	}
}

func (f *allFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, loop interpreter.Looper) (interface{}, error) {
	loopEvaluator := loop.Evaluator()
	col := wrapCollection(ctx, node)
	count := col.Count()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if b, ok := res.(interpreter.BooleanAccessor); !ok {
			return nil, fmt.Errorf("Expected boolean result from loop, got %T", res)
		} else if !b.Bool() {
			return interpreter.False, nil
		}
	}

	return interpreter.True, nil
}

type allAnyTrueFalseFunction struct {
	interpreter.BaseFunction
	all bool
	t   bool
}

func newAllAnyTrueFalseFunction(name string, all, t bool) *allAnyTrueFalseFunction {
	return &allAnyTrueFalseFunction{
		interpreter.NewBaseFunction(name, -1, 0, 0),
		all,
		t,
	}
}

func newAllTrueFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("allTrue", true, true)
}

func newAllFalseFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("allFalse", true, false)
}

func newAnyTrueFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("anyTrue", false, true)
}

func newAnyFalseFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("anyFalse", false, false)
}

func (f *allAnyTrueFalseFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()

	for i := 0; i < count; i++ {
		this := col.Get(i)
		if b, ok := this.(interpreter.BooleanAccessor); !ok {
			return nil, fmt.Errorf("Expected boolean result from loop, got %T", this)
		} else if f.all && f.t != b.Bool() {
			return interpreter.False, nil
		} else if !f.all && f.t == b.Bool() {
			return interpreter.True, nil
		}
	}

	if f.all {
		return interpreter.True, nil
	} else {
		return interpreter.False, nil
	}
}

type subsetOfFunction struct {
	interpreter.BaseFunction
}

func newSubsetOfFunction() *subsetOfFunction {
	return &subsetOfFunction{
		interpreter.NewBaseFunction("subsetOf", -1, 1, 1),
	}
}

func (f *subsetOfFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count > 0 {
		otherCol := wrapCollection(ctx, args[0])
		for i := 0; i < count; i++ {
			if !otherCol.Contains(col.Get(i)) {
				return interpreter.False, nil
			}
		}
	}

	return interpreter.True, nil
}

type supersetOfFunction struct {
	interpreter.BaseFunction
}

func newSupersetOfFunction() *supersetOfFunction {
	return &supersetOfFunction{
		interpreter.NewBaseFunction("supersetOf", -1, 1, 1),
	}
}

func (f *supersetOfFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	otherCol := wrapCollection(ctx, args[0])
	count := otherCol.Count()
	if count > 0 {
		col := wrapCollection(ctx, node)
		for i := 0; i < count; i++ {
			if !col.Contains(otherCol.Get(i)) {
				return interpreter.False, nil
			}
		}
	}

	return interpreter.True, nil
}

type countFunction struct {
	interpreter.BaseFunction
}

func newCountFunction() *countFunction {
	return &countFunction{
		interpreter.NewBaseFunction("count", -1, 0, 0),
	}
}

func (f *countFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	return interpreter.NewInteger(int32(col.Count())), nil
}

type distinctFunction struct {
	interpreter.BaseFunction
}

func newDistinctFunction() *distinctFunction {
	return &distinctFunction{
		interpreter.NewBaseFunction("distinct", -1, 0, 0),
	}
}

func (f *distinctFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Count() < 2 {
		return col, nil
	}

	res := ctx.NewCol()
	res.AddAllUnique(col)
	return res, nil
}

type isDistinctFunction struct {
	interpreter.BaseFunction
}

func newIsDistinctFunction() *isDistinctFunction {
	return &isDistinctFunction{
		BaseFunction: interpreter.NewBaseFunction("isDistinct", -1, 0, 0),
	}
}

func (f *isDistinctFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}

	if col.Count() == 1 {
		return interpreter.True, nil
	}

	res := ctx.NewCol()
	res.AddAllUnique(col)
	return interpreter.BooleanOf(col.Count() == res.Count()), nil
}
