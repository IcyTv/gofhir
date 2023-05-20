package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

var functions = []interpreter.FunctionExecutor{
	// existence
	newEmptyFunction(),
	newExistsFunction(),
	newAllFunction(),
	newAllTrueFunction(),
	newAnyTrueFunction(),
	newAnyFalseFunction(),
	newAllFalseFunction(),
	newSubsetOfFunction(),
	newSupersetOfFunction(),
	newCountFunction(),
	newDistinctFunction(),
	newIsDistinctFunction(),
	// Filtering and projection
	newWhereFunction(),
	newSelectFunction(),
	repeatFunc,
	newOfTypeFunction(),
	// Sub-Setting
	newSingleFunction(),
	newFirstFunction(),
	newLastFunction(),
	newTailFunction(),
	newSkipFunction(),
	newTakeFunction(),
	newIntersectFunction(),
	newExcludeFunction(),
	// Combining
	newUnionFunction(),
	newCombineFunction(),
	// Conversion
	newIIfFunction(),
	toBooleanFunc,
	newConvertsToBooleanFunction(),
	toIntegerFunc,
	newConvertsToIntegerFunction(),
	toDateFunc,
	newConvertsToDateFunction(),
	toDateTimeFunc,
	newConvertsToDateTimeFunction(),
	toDecimalFunc,
	newConvertsToDecimalFunction(),
	toQuantityFunc,
	newConvertsToQuantityFunction(),
	toStringFunc,
	newConvertsToStringFunction(),
	toTimeFunc,
	newConvertsToTimeFunction(),
	// String manipulation
	newIndexOfFunction(),
	newSubstringFunction(),
	newStartsWithFunction(),
	newEndsWithFunction(),
	newContainsFunction(),
	newUpperFunction(),
	newLowerFunction(),
	newReplaceFunction(),
	newMatchesFunction(),
	newReplaceMatchesFunction(),
	newLengthFunction(),
	newToCharsFunction(),
	// Math
	newAbsFunction(),
	newCeilingFunction(),
	newExpFunction(),
	newFloorFunction(),
	newLnFunction(),
	newLogFunction(),
	newPowerFunction(),
	newRoundFunction(),
	newSqrtFunction(),
	newTruncateFunction(),
	// tree navigation
	childrenFunc,
	newDescendantsFunction(),
	// utility
	newTraceFunction(),
	newNowFunction(),
	newTimeOfDayFunction(),
	newTodayFunction(),
	// type
	newAsFunction(),
	newIsFunction(),
	// Aggregate
	newAggregateFunction(),
}

var functionsByName = createFunctionsByName(functions)

type FunctionInvocation struct {
	executor        interpreter.FunctionExecutor
	paramEvaluators []interpreter.Evaluator
}

func LookupFunctionInvocation(name string, paramEvaluators []interpreter.Evaluator) (*FunctionInvocation, error) {
	executor, found := functionsByName[name]
	if !found {
		return nil, fmt.Errorf("Function '%s' not found", name)
	}

	if len(paramEvaluators) < executor.MinParams() {
		return nil, fmt.Errorf("Function '%s' requires at least %d parameters", name, executor.MinParams())
	}
	if len(paramEvaluators) > executor.MaxParams() {
		return nil, fmt.Errorf("Function '%s' requires at most %d parameters", name, executor.MaxParams())
	}

	return newFunctionInvocation(executor, paramEvaluators), nil
}

func newFunctionInvocation(executor interpreter.FunctionExecutor, paramEvaluators []interpreter.Evaluator) *FunctionInvocation {
	return &FunctionInvocation{executor, paramEvaluators}
}

func (f *FunctionInvocation) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	var args []interface{}
	ac := len(f.paramEvaluators)
	if ac == 0 {
		args = nil
	} else {
		evaluatorParam := f.executor.EvaluatorParam()
		args = make([]interface{}, ac)

		var loopEvaluator interpreter.Evaluator
		for pos, argEvaluator := range f.paramEvaluators {
			if evaluatorParam == pos {
				loopEvaluator = argEvaluator
			} else {
				if argEvaluator != nil {
					if arg, err := argEvaluator.Evaluate(ctx, node, loop); err != nil {
						return nil, fmt.Errorf("Error in argument %d of executor invocation %s: %v", pos, f.executor.Name(), err)
					} else {
						args[pos] = arg
					}
				}
			}
		}

		if evaluatorParam >= 0 {
			loop = interpreter.NewLoop(loopEvaluator)
		}
	}

	return f.executor.Execute(ctx, node, args, loop)
}

func createFunctionsByName(functions []interpreter.FunctionExecutor) map[string]interpreter.FunctionExecutor {
	functionsByName := make(map[string]interpreter.FunctionExecutor, len(functions))
	for _, f := range functions {
		functionsByName[f.Name()] = f
	}
	return functionsByName
}
