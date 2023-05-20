package expression

import (
	"fmt"

	"github.com/gofhir/fhirpath/interpreter"
)

type MemberInvocation struct {
	name string
}

func NewMemberInvocation(name string) *MemberInvocation {
	return &MemberInvocation{name}
}

func (i *MemberInvocation) Evaluate(ctx interpreter.ContextAccessor, node interface{}, loop interpreter.Looper) (interface{}, error) {
	if node == nil {
		return nil, fmt.Errorf("Cannot extract path from empty: %s", i.name)
	}

	return ctx.ModelAdapter().Navigate(node, i.name)
}
