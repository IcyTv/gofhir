package visitors

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitParenthesizedTerm(ctx *parser.ParenthesizedTermContext) interface{} {
	return v.VisitChild(ctx, 1)
}

func (v *Visitor) VisitLiteralTerm(ctx *parser.LiteralTermContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitExternalConstantTerm(ctx *parser.ExternalConstantTermContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitExternalConstant(ctx *parser.ExternalConstantContext) interface{} {
	return v.visitTree(ctx, 2, visitExternalConstant)
}

func visitExternalConstant(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	name := args[1].(string)
	return expression.ParseExtConstantTerm(expression.ExtractIdentifier(name)), nil
}

func (v *Visitor) VisitInvocationTerm(ctx *parser.InvocationTermContext) interface{} {
	return v.visitTree(ctx, 1, visitInvocationTerm)
}

func visitInvocationTerm(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	invocationEvaluator := args[0].(interpreter.Evaluator)
	return expression.NewInvocationTerm(invocationEvaluator), nil
}
