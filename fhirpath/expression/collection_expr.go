package expression

import "github.com/gofhir/fhirpath/interpreter"

type CollectionExpression struct {
	eval interpreter.Evaluator
}

func NewCollectionExpression(eval interpreter.Evaluator) CollectionExpression {
	return CollectionExpression{eval}
}

func (e *CollectionExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	res, err := e.eval.Evaluate(ctx, node, loop)
	if err != nil {
		return res, err
	}

	if res == nil {
		return interpreter.EmptyColl, nil
	}
	if c, ok := res.(interpreter.ColAccessor); ok {
		return c, nil
	} else {
		return ctx.NewColWithItem(res), nil
	}
}
