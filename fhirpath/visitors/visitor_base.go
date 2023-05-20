package visitors

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/parser"
)

type Visitor struct {
	parser.BaseFHIRPathVisitor
	errorItemCollection *ErrorItemCollection
}

type visitorFunc func(ctx antlr.ParserRuleContext) (interpreter.Evaluator, error)
type visitorArgFunc func(ctx antlr.ParserRuleContext, args []interface{}) (interpreter.Evaluator, error)

func NewVisitor(errorItemCollection *ErrorItemCollection) *Visitor {
	v := new(Visitor)
	v.errorItemCollection = errorItemCollection
	return v
}

func (v *Visitor) AddError(ctx antlr.ParserRuleContext, msg string) interpreter.Evaluator {
	start := ctx.GetStart()
	v.errorItemCollection.AddError(start.GetLine(), start.GetColumn(), msg)
	return nil
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	count := node.GetChildCount()
	if count == 0 {
		return nil
	}

	r := make([]interface{}, count)
	for pos, child := range node.GetChildren() {
		n := v.evalNode(child)
		if n == nil {
			return nil
		}
		r[pos] = n
	}

	return r
}

func (v *Visitor) VisitFirstChild(node antlr.RuleNode) interface{} {
	return v.VisitChild(node, 0)
}

func (v *Visitor) VisitChild(node antlr.RuleNode, i int) interface{} {
	if node.GetChildCount() <= i {
		return nil
	}
	return v.evalNode(node.GetChild(i))
}

func (v *Visitor) evalNode(node antlr.Tree) interface{} {
	switch n := node.(type) {
	case antlr.RuleNode:
		return n.Accept(v)
	case antlr.ErrorNode:
		return nil
	case antlr.TerminalNode:
		return n.GetText()
	}
	return nil
}

func (v *Visitor) visit(ctx antlr.ParserRuleContext, f visitorFunc) interface{} {
	if l, err := f(ctx); err != nil {
		return v.AddError(ctx, err.Error())
	} else {
		return l
	}
}

func (v *Visitor) visitTree(ctx antlr.ParserRuleContext, argCount int, f visitorArgFunc) interpreter.Evaluator {
	c := v.VisitChildren(ctx)

	args, ok := c.([]interface{})
	if !ok || len(args) < argCount {
		return nil
	}

	if l, err := f(ctx, args); err != nil {
		return v.AddError(ctx, err.Error())
	} else {
		return l
	}
}
