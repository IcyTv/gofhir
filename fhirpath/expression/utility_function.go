package expression

import (
	"time"

	"github.com/gofhir/fhirpath/interpreter"
)

type traceFunction struct {
	interpreter.BaseFunction
}

func newTraceFunction() *traceFunction {
	return &traceFunction{
		interpreter.NewBaseFunction("trace", 1, 1, 2),
	}
}

func (f *traceFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, args []interface{}, loop interpreter.Looper) (interface{}, error) {
	tracer := ctx.Tracer()
	if tracer == nil {
		return node, nil
	}

	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return node, err
	}

	col := wrapCollection(ctx, node)
	count := col.Count()

	var traced interpreter.ColAccessor
	if count == 0 {
		traced = interpreter.EmptyColl
	} else if loopEvaluator := loop.Evaluator(); loopEvaluator != nil {
		projected := ctx.NewCol()
		for i := 0; i < count; i++ {
			this := col.Get(i)
			loop.IncIndex(this)

			res, err := loopEvaluator.Evaluate(ctx, this, loop)
			if err != nil {
				return nil, err
			}
			projected.Add(res)
		}
		traced = projected
	} else {
		traced = col
	}

	tracer.Trace(name.String(), traced)
	return node, nil
}

type nowFunction struct {
	interpreter.BaseFunction
}

func newNowFunction() *nowFunction {
	return &nowFunction{
		interpreter.NewBaseFunction("now", -1, 0, 0),
	}
}

func (f *nowFunction) Execute(interpreter.ContextAccessor, interface{}, []interface{}, interpreter.Looper) (interface{}, error) {
	return interpreter.NewDateTime(time.Now()), nil
}

type timeOfDayFunction struct {
	interpreter.BaseFunction
}

func newTimeOfDayFunction() *timeOfDayFunction {
	return &timeOfDayFunction{
		interpreter.NewBaseFunction("timeOfDay", -1, 0, 0),
	}
}

func (f *timeOfDayFunction) Execute(interpreter.ContextAccessor, interface{}, []interface{}, interpreter.Looper) (interface{}, error) {
	return interpreter.NewTime(time.Now()), nil
}

type todayFunction struct {
	interpreter.BaseFunction
}

func newTodayFunction() *todayFunction {
	return &todayFunction{
		interpreter.NewBaseFunction("today", -1, 0, 0),
	}
}

func (f *todayFunction) Execute(interpreter.ContextAccessor, interface{}, []interface{}, interpreter.Looper) (interface{}, error) {
	return interpreter.NewDate(time.Now()), nil
}
