package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type ThisInvocation struct{}

var thisInvocation = &ThisInvocation{}

func NewThisInvocation() *ThisInvocation {
	return thisInvocation
}

func (t *ThisInvocation) Evaluate(_ interpreter.ContextAccessor, _ interface{}, loop interpreter.Looper) (interface{}, error) {
	if loop == nil {
		return nil, fmt.Errorf("This invocation is not in a loop")
	}
	return loop.This(), nil
}

type IndexInvocation struct{}

var indexInvocation = &IndexInvocation{}

func NewIndexInvocation() *IndexInvocation {
	return indexInvocation
}

func (t *IndexInvocation) Evaluate(_ interpreter.ContextAccessor, _ interface{}, loop interpreter.Looper) (interface{}, error) {
	if loop == nil {
		return nil, fmt.Errorf("Index invocation is not in a loop")
	}

	return interpreter.NewInteger(int32(loop.Index())), nil
}

type TotalInvocation struct{}

var totalInvocation = &TotalInvocation{}

func NewTotalInvocation() *TotalInvocation {
	return totalInvocation
}

func (t *TotalInvocation) Evaluate(_ interpreter.ContextAccessor, _ interface{}, loop interpreter.Looper) (interface{}, error) {
	if loop == nil {
		return nil, fmt.Errorf("Total invocation is not in a loop")
	}

	return loop.Total(), nil
}
