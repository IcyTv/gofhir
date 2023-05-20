package visitors

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func visitArithmeticExpression(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	leftEvaluator := args[0].(interpreter.Evaluator)
	stringOp := args[1].(string)
	rightEvaluator := args[2].(interpreter.Evaluator)

	var op interpreter.ArithmeticOps
	switch stringOp {
	case "+":
		op = interpreter.AdditionOp
	case "-":
		op = interpreter.SubtractionOp
	case "*":
		op = interpreter.MultiplicationOp
	case "/":
		op = interpreter.DivisionOp
	case "div":
		op = interpreter.DivOp
	case "mod":
		op = interpreter.ModOp
	case "&":
		return expression.NewStringConcatExpression(leftEvaluator, rightEvaluator), nil
	default:
		return nil, fmt.Errorf("Unsupported arithmetic operator: %s", stringOp)
	}

	return expression.NewArithmeticExpression(leftEvaluator, op, rightEvaluator), nil
}
