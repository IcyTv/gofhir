package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type singleFunction struct {
	interpreter.BaseFunction
}

func newSingleFunction() *singleFunction {
	return &singleFunction{
		interpreter.NewBaseFunction("single", -1, 0, 0),
	}
}

func (f *singleFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if count > 0 {
		return nil, fmt.Errorf("Expected collection with one item, had %d", count)
	}
	return col.Get(0), nil
}

type firstFunction struct {
	interpreter.BaseFunction
}

func newFirstFunction() *firstFunction {
	return &firstFunction{
		interpreter.NewBaseFunction("first", -1, 0, 0),
	}
}

func (f *firstFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}
	return col.Get(0), nil
}

type lastFunction struct {
	interpreter.BaseFunction
}

func newLastFunction() *lastFunction {
	return &lastFunction{
		interpreter.NewBaseFunction("last", -1, 0, 0),
	}
}

func (f *lastFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}
	return col.Get(col.Count() - 1), nil
}

type tailFunction struct {
	interpreter.BaseFunction
}

func newTailFunction() *tailFunction {
	return &tailFunction{
		interpreter.NewBaseFunction("tail", -1, 0, 0),
	}
}

func (f *tailFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count < 2 {
		return nil, nil
	}

	tail := ctx.NewCol()
	for i := 1; i < count; i++ {
		tail.Add(col.Get(i))
	}

	return tail, nil
}

type skipFunction struct {
	interpreter.BaseFunction
}

func newSkipFunction() *skipFunction {
	return &skipFunction{
		interpreter.NewBaseFunction("skip", -1, 1, 1),
	}
}

func (f *skipFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	var num int
	if n, ok := unwrapCollection(args[0]).(interpreter.NumberAccessor); !ok {
		return nil, fmt.Errorf("Expected integer, got %T", args[0])
	} else {
		num = int(n.Int())
	}

	col := wrapCollection(ctx, node)
	if num <= 0 {
		return col, nil
	}

	count := col.Count()
	if count <= num {
		return nil, nil
	}

	res := ctx.NewCol()
	for i := num; i < count; i++ {
		res.Add(col.Get(i))
	}

	return res, nil
}

type takeFunction struct {
	interpreter.BaseFunction
}

func newTakeFunction() *takeFunction {
	return &takeFunction{
		interpreter.NewBaseFunction("take", -1, 1, 1),
	}
}

func (f *takeFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	var num int
	if n, ok := unwrapCollection(args[0]).(interpreter.NumberAccessor); !ok {
		return nil, fmt.Errorf("Expected integer, got %T", args[0])
	} else {
		num = int(n.Int())
	}
	if num <= 0 {
		return nil, nil
	}

	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if count <= num {
		return col, nil
	}

	res := ctx.NewCol()
	for i := 0; i < num; i++ {
		res.Add(col.Get(i))
	}
	return res, nil
}

type intersectFunction struct {
	interpreter.BaseFunction
}

func newIntersectFunction() *intersectFunction {
	return &intersectFunction{
		interpreter.NewBaseFunction("intersect", -1, 1, 1),
	}
}

func (f *intersectFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	other := wrapCollection(ctx, args[0])
	if other.Empty() {
		return nil, nil
	}

	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}

	if col.Count() > other.Count() {
		x := col
		col = other
		other = x
	}
	count := col.Count()

	res := ctx.NewCol()
	for i := 0; i < count; i++ {
		n := col.Get(i)
		if other.Contains(n) {
			res.AddUnique(n)
		}
	}

	return res, nil
}

type excludeFunction struct {
	interpreter.BaseFunction
}

func newExcludeFunction() *excludeFunction {
	return &excludeFunction{
		interpreter.NewBaseFunction("exclude", -1, 1, 1),
	}
}

func (f *excludeFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	other := wrapCollection(ctx, args[0])
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if other.Empty() {
		return col, nil
	}

	res := ctx.NewCol()
	for i := 0; i < count; i++ {
		n := col.Get(i)
		if !other.Contains(n) {
			res.Add(n)
		}
	}

	return res, nil
}
