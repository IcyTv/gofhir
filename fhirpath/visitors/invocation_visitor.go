package visitors

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/expression"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

func (v *Visitor) VisitFunctionInvocation(ctx *parser.FunctionInvocationContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	return v.visitTree(ctx, 3, visitFunction)
}

func visitFunction(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	name := args[0].(string)

	var paramEval []interpreter.Evaluator
	if len(args) < 4 {
		paramEval = []interpreter.Evaluator{}
	} else if name == "as" || name == "is" {
		typeSpec := args[2].(string)
		paramEval = []interpreter.Evaluator{expression.NewRawStringLiteral(typeSpec)}
	} else {
		paramList := args[2].([]interface{})
		paramEval = make([]interpreter.Evaluator, (len(paramList)+1)/2)
		for pos, param := range paramList {
			if pos%2 == 0 {
				paramEval[pos/2] = param.(interpreter.Evaluator)
			}
		}
	}

	return expression.LookupFunctionInvocation(expression.ExtractIdentifier(name), paramEval)
}

func (v *Visitor) VisitParamList(ctx *parser.ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitThisInvocation(*parser.ThisInvocationContext) interface{} {
	return expression.NewThisInvocation()
}

func (v *Visitor) VisitIndexInvocation(*parser.IndexInvocationContext) interface{} {
	return expression.NewIndexInvocation()
}

func (v *Visitor) VisitTotalInvocation(*parser.TotalInvocationContext) interface{} {
	return expression.NewTotalInvocation()
}

func (v *Visitor) VisitMemberInvocation(ctx *parser.MemberInvocationContext) interface{} {
	return v.visitTree(ctx, 1, visitMemberInvocation)
}

func visitMemberInvocation(_ antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error) {
	name := args[0].(string)
	return expression.NewMemberInvocation(expression.ExtractIdentifier(name)), nil
}
