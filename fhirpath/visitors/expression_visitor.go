package visitors

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitTermExpression(ctx *parser.TermExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitPolarityExpression(ctx *parser.PolarityExpressionContext) interface{} {
	return v.visitTree(ctx, 2, visitPolarityExpression)
}

func visitPolarityExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	op := args[0].(string)
	evaluator := args[1].(interpreter.Evaluator)

	if op == "-" && evaluator != nil {
		evaluator = expression.NewNegatorExpression(evaluator.(interpreter.Evaluator))
	}

	return evaluator, nil
}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	log.Println("VisitEqualityExpression", ctx.GetText())
	return v.visitTree(ctx, 3, visitEqualityExpression)
}

func visitEqualityExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	evalLeft := args[0].(interpreter.Evaluator)
	op := args[1].(string)
	evalRight := args[2].(interpreter.Evaluator)

	not := false
	equivalent := false
	switch op {
	case "=":
	case "!=":
		not = true
	case "~":
		equivalent = true
	case "!~":
		not = true
		equivalent = true
	default:
		return nil, fmt.Errorf("invalid equality operator: %s", op)
	}

	return expression.NewEqualityExpression(not, equivalent, evalLeft, evalRight), nil
}

func (v *Visitor) VisitUnionExpression(ctx *parser.UnionExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitUnionExpression)
}

func visitUnionExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	left := args[0].(interpreter.Evaluator)
	right := args[2].(interpreter.Evaluator)

	return expression.NewUnionExpression(left, right), nil
}

func (v *Visitor) VisitIndexerExpression(ctx *parser.IndexerExpressionContext) interface{} {
	return v.visitTree(ctx, 4, visitIndexerExpression)
}

func visitIndexerExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	expr := args[0].(interpreter.Evaluator)
	index := args[2].(interpreter.Evaluator)

	return expression.NewIndexerExpression(expr, index), nil
}

func (v *Visitor) VisitInvocationExpression(ctx *parser.InvocationExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitInvocationExpression)
}

func visitInvocationExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	exprEvaluator := args[0].(interpreter.Evaluator)
	invocationEvaluator := args[2].(interpreter.Evaluator)

	return expression.NewInvocationExpression(exprEvaluator, invocationEvaluator), nil
}

func (v *Visitor) VisitInequalityExpression(ctx *parser.InequalityExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitInequalityExpression)
}

func visitInequalityExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	evalLeft := args[0].(interpreter.Evaluator)
	op := args[1].(string)
	evalRight := args[2].(interpreter.Evaluator)

	var cmpOp expression.ComparisonOp
	switch op {
	case "<=":
		cmpOp = expression.LessOrEqualThanOp
	case "<":
		cmpOp = expression.LessThanOp
	case ">=":
		cmpOp = expression.GreaterOrEqualThanOp
	case ">":
		cmpOp = expression.GreaterThanOp
	default:
		return nil, fmt.Errorf("invalid inequality operator: %s", op)
	}

	return expression.NewComparisonExpression(evalLeft, cmpOp, evalRight), nil
}

func (v *Visitor) VisitMembershipExpression(ctx *parser.MembershipExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitMembershipExpression)
}

func visitMembershipExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	left := args[0].(interpreter.Evaluator)
	op := args[1].(string)
	right := args[2].(interpreter.Evaluator)

	switch op {
	case "in":
		return expression.NewContainsExpression(left, right, true), nil
	case "contains":
		return expression.NewContainsExpression(left, right, false), nil
	default:
		return nil, fmt.Errorf("invalid membership operator: %s", op)
	}
}

func (v *Visitor) VisitAndExpression(ctx *parser.AndExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitBooleanExpression)
}

func (v *Visitor) VisitOrExpression(ctx *parser.OrExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitBooleanExpression)
}

func (v *Visitor) VisitImpliesExpression(ctx *parser.ImpliesExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitBooleanExpression)
}

func visitBooleanExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	left := args[0].(interpreter.Evaluator)
	op := args[1].(string)
	right := args[2].(interpreter.Evaluator)

	var booleanOp expression.BooleanOp
	switch op {
	case "and":
		booleanOp = expression.AndOp
	case "or":
		booleanOp = expression.OrOp
	case "xor":
		booleanOp = expression.XorOp
	case "implies":
		booleanOp = expression.ImpliesOp
	default:
		return nil, fmt.Errorf("invalid boolean operator: %s", op)
	}

	return expression.NewBooleanExpression(left, booleanOp, right), nil
}

func (v *Visitor) VisitTypeExpression(ctx *parser.TypeExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitTypeExpression)
}

func visitTypeExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	eval := args[0].(interpreter.Evaluator)
	op := args[1].(string)
	name := expression.ExtractIdentifier(args[2].(string))

	switch op {
	case "as":
		return expression.NewAsTypeExpression(eval, name)
	case "is":
		return expression.NewIsTypeExpression(eval, name)
	default:
		return nil, fmt.Errorf("invalid type operator: %s", op)
	}
}
