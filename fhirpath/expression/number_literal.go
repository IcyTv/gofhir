package expression

import (
	"strings"

	"github.com/gofhir/fhirpath/interpreter"
)

type NumberLiteral struct {
	node interpreter.NumberAccessor
}

func NewNumberLiteralInt(value int32) interpreter.Evaluator {
	return &NumberLiteral{interpreter.NewInteger(value)}
}

func NewNumberLiteralFloat64(value float64) interpreter.Evaluator {
	return &NumberLiteral{interpreter.NewDecimalFloat64(value)}
}

func ParseNumberLiteral(value string) (interpreter.Evaluator, error) {
	var node interpreter.NumberAccessor
	var err error
	if strings.ContainsRune(value, '.') {
		node, err = interpreter.ParseDecimal(value)
	} else {
		node, err = interpreter.ParseInteger(value)
	}

	if err != nil {
		return nil, err
	} else {
		return &NumberLiteral{node}, nil
	}
}

func (e *NumberLiteral) Evaluate(ctx interpreter.ContextAccessor, _ interface{}, _ interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
