package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type IndexerExpression struct {
	exprEvaluator  interpreter.Evaluator
	indexEvaluator interpreter.Evaluator
}

func NewIndexerExpression(exprEvaluator, indexEvaluator interpreter.Evaluator) *IndexerExpression {
	return &IndexerExpression{exprEvaluator, indexEvaluator}
}

func (e *IndexerExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, looper interpreter.Looper) (interface{}, error) {
	col, err := e.exprEvaluator.Evaluate(ctx, node, looper)
	if err != nil {
		return nil, err
	}
	index, err := e.indexEvaluator.Evaluate(ctx, node, looper)
	if err != nil {
		return nil, err
	}

	if col == nil || index == nil {
		return nil, nil
	}

	var indexValue int
	if n, ok := index.(interpreter.NumberAccessor); !ok {
		return nil, fmt.Errorf("Index is not a number %T", index)
	} else {
		indexValue = int(n.Int())
	}

	if indexValue < 0 {
		return nil, nil
	}

	if c, ok := col.(interpreter.ColAccessor); ok {
		if indexValue >= c.Count() {
			return nil, nil
		}
		return c.Get(indexValue), nil
	}

	if indexValue > 0 {
		return nil, nil
	}

	return col, nil
}
