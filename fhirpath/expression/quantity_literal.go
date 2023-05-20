package expression

import (
	"fmt"
	"regexp"

	"github.com/gofhir/fhirpath/interpreter"
)

var quantityUnitRegexp = regexp.MustCompile("^[^\\s]+(\\s[^\\s]+)*$")

type QuantityLiteral struct {
	node interpreter.QuantityAccessor
}

func ParseQuantityLiteral(number string, unit string) (interpreter.Evaluator, error) {
	value, err := interpreter.ParseDecimal(number)
	if err != nil {
		return nil, err
	}
	convertedUnit, err := parseQuantityUnit(unit)
	if err != nil {
		return nil, err
	}

	return &QuantityLiteral{interpreter.NewQuantity(value, convertedUnit)}, nil
}

func parseQuantityUnit(unit string) (interpreter.StringAccessor, error) {
	if len(unit) == 0 || unit == "''" {
		return nil, nil
	}

	u := parseStringLiteral(unit, stringDelimiterChar)
	if !quantityUnitRegexp.MatchString(u) {
		return nil, fmt.Errorf("Invalid quantity unit '%s'", u)
	}

	return interpreter.NewString(u), nil
}

func (e *QuantityLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
