package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type DateTimeLiteral struct {
	node interpreter.DateTimeAccessor
}

func ParseDateTimeLiteral(value string) (interpreter.Evaluator, error) {
	if len(value) < 2 || value[0] != '@' {
		return nil, fmt.Errorf("Invalid date literal: %s", value)
	}
	if node, err := interpreter.ParseDateTime(value[1:]); err != nil {
		return nil, err
	} else {
		return &DateTimeLiteral{node}, nil
	}
}

func (e *DateTimeLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
