package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type DateLiteral struct {
	node interpreter.DateAccessor
}

func ParseDateLiteral(value string) (interpreter.Evaluator, error) {
	if len(value) < 2 || value[0] != '@' {
		return nil, fmt.Errorf("Invalid date literal: %s", value)
	}
	if node, err := interpreter.ParseDate(value[1:]); err != nil {
		return nil, err
	} else {
		return &DateLiteral{node}, nil
	}
}

func (e *DateLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
