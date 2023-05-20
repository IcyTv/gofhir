package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type ExtConstantTerm struct {
	name string
}

func ParseExtConstantTerm(value string) *ExtConstantTerm {
	return &ExtConstantTerm{value}
}

func (e *ExtConstantTerm) Evaluate(ctx interpreter.ContextAccessor, _ interface{}, _ interpreter.Looper) (interface{}, error) {
	res, found := ctx.EnvVar(e.name)
	if !found {
		return nil, fmt.Errorf("Environment variable '%s' not found", e.name)
	}

	return res, nil
}
