package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type TimeLiteral struct {
	node interpreter.TimeAccessor
}

func ParseTimeLiteral(value string) (interpreter.Evaluator, error) {
	if len(value) < 3 || value[0] != '@' || value[1] != 'T' {
		return nil, fmt.Errorf("Invalid time literal: %s", value)
	}
	if node, err := interpreter.ParseTime(value[2:]); err != nil {
		return nil, err
	} else {
		return &TimeLiteral{node}, nil
	}
}

func (e *TimeLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
