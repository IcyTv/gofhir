package expression

import "github.com/gofhir/fhirpath/interpreter"

type BooleanLiteral struct {
	node interpreter.BooleanAccessor
}

func NewBooleanLiteral(value bool) interpreter.Evaluator {
	return &BooleanLiteral{interpreter.BooleanOf(value)}
}

func ParseBooleanLiteral(value string) (interpreter.Evaluator, error) {
	if node, err := interpreter.ParseBoolean(value); err != nil {
		return nil, err
	} else {
		return &BooleanLiteral{node}, nil
	}
}

func (e *BooleanLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
