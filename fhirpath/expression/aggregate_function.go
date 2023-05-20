package expression

import "github.com/gofhir/fhirpath/interpreter"

type aggregateFunction struct {
	interpreter.BaseFunction
}

func newAggregateFunction() *aggregateFunction {
	return &aggregateFunction{
		BaseFunction: interpreter.NewBaseFunction("aggregate", 0, 1, 2),
	}
}

func (f *aggregateFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, loop interpreter.Looper) (interface{}, error) {
	if node == nil {
		return interpreter.False, nil
	}

	if len(args) > 1 {
		loop.SetTotal(args[1])
	}

	loopEvaluator := loop.Evaluator()
	if loopEvaluator != nil {
		col := wrapCollection(ctx, node)
		count := col.Count()
		for i := 0; i < count; i++ {
			this := col.Get(i)
			loop.IncIndex(this)

			res, err := loopEvaluator.Evaluate(ctx, this, loop)
			if err != nil {
				return nil, err
			}
			loop.SetTotal(res)
		}
	}

	return loop.Total(), nil
}
