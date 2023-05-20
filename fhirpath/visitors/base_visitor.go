package visitors

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) interface{} {
	return v.visitQualifiedIdentifier(ctx)
}

func (v *Visitor) visitQualifiedIdentifier(node antlr.RuleNode) interface{} {
	c := v.VisitChildren(node)

	parts, ok := c.([]interface{})
	if !ok {
		return nil
	}

	var b strings.Builder
	b.Grow(16)

	for pos, part := range parts {
		if pos%2 == 0 {
			b.WriteString(expression.ExtractIdentifier(part.(string)))
		} else {
			b.WriteByte('.')
		}
	}

	return b.String()
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return v.VisitFirstChild(ctx)
}
