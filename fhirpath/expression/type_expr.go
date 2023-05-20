package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type AsTypeExpression struct {
	exprEvaluator interpreter.Evaluator
	fqName        interpreter.FQTypeNameAccessor
}

func NewAsTypeExpression(exprEvaluator interpreter.Evaluator, name string) (*AsTypeExpression, error) {
	fqName, err := interpreter.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}
	return &AsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *AsTypeExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(interpreter.ColAccessor); ok {
		return nil, fmt.Errorf("as operator cannot be applied on a collection")
	}

	return interpreter.CastModelType(ctx.ModelAdapter(), item, e.fqName)
}

type IsTypeExpression struct {
	exprEvaluator interpreter.Evaluator
	fqName        interpreter.FQTypeNameAccessor
}

func NewIsTypeExpression(exprEvaluator interpreter.Evaluator, name string) (*IsTypeExpression, error) {
	fqName, err := interpreter.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}
	return &IsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *IsTypeExpression) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(interpreter.ColAccessor); ok {
		return nil, fmt.Errorf("is operator cannot be applied on a collection")
	}

	return interpreter.BooleanOf(interpreter.HasModelType(ctx.ModelAdapter(), item, e.fqName)), nil
}
