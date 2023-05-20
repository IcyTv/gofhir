package visitors

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitNullLiteral(*parser.NullLiteralContext) interface{} {
	return expression.NewEmptyLiteral()
}

func (v *Visitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	return v.visit(ctx, visitBooleanLiteral)
}

func visitBooleanLiteral(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error) {
	return expression.ParseBooleanLiteral(ctx.GetText())
}

func (v *Visitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return expression.ParseStringLiteral(ctx.GetText())
}

func (v *Visitor) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	return v.visit(ctx, visitNumberLiteral)
}

func visitNumberLiteral(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error) {
	return expression.ParseNumberLiteral(ctx.GetText())
}

func (v *Visitor) VisitDateLiteral(ctx *parser.DateLiteralContext) interface{} {
	return v.visit(ctx, visitDateLiteral)
}

func visitDateLiteral(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error) {
	return expression.ParseDateLiteral(ctx.GetText())
}

func (v *Visitor) VisitDateTimeLiteral(ctx *parser.DateTimeLiteralContext) interface{} {
	return v.visit(ctx, visitDateTimeLiteral)
}

func visitDateTimeLiteral(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error) {
	return expression.ParseDateTimeLiteral(ctx.GetText())
}

func (v *Visitor) VisitTimeLiteral(ctx *parser.TimeLiteralContext) interface{} {
	return v.visit(ctx, visitTimeLiteral)
}

func visitTimeLiteral(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error) {
	return expression.ParseTimeLiteral(ctx.GetText())
}

func (v *Visitor) VisitQuantity(ctx *parser.QuantityContext) interface{} {
	return v.visitTree(ctx, 2, visitQuantity)
}

func visitQuantity(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	number := args[0].(string)
	unit := args[1].(string)
	return expression.ParseQuantityLiteral(number, unit)
}

func (v *Visitor) VisitQuantityLiteral(ctx *parser.QuantityLiteralContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitUnit(ctx *parser.UnitContext) interface{} {
	return ctx.GetText()
}
