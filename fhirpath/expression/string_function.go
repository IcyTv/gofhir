package expression

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/gofhir/fhirpath/interpreter"
)

type indexOfFunction struct {
	interpreter.BaseFunction
}

func newIndexOfFunction() *indexOfFunction {
	return &indexOfFunction{
		interpreter.NewBaseFunction("indexOf", -1, 1, 1),
	}
}

func (f *indexOfFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}
	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	i := strings.Index(s.String(), ss.String())
	return interpreter.NewInteger(int32(i)), nil
}

type substringFunction struct {
	interpreter.BaseFunction
}

func newSubstringFunction() *substringFunction {
	return &substringFunction{
		interpreter.NewBaseFunction("substring", -1, 1, 2),
	}
}

func (f *substringFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	start, err := integerNode(args[0])
	if start == nil || err != nil {
		return nil, err
	}
	startVal := start.Int()
	if startVal < 0 {
		startVal = 0
	}

	var l interpreter.IntegerAccessor = nil
	if len(args) > 1 {
		l, err = integerNode(args[1])
		if l == nil || err != nil {
			return nil, err
		}
	}

	var lVal int32
	if l != nil {
		lVal = l.Int()
		if lVal <= 0 {
			return nil, nil
		}
	}

	sr := []rune(s.String())
	srLen := int32(len(sr))
	if l == nil {
		lVal = srLen - startVal
	}

	if startVal >= srLen || lVal <= 0 {
		return nil, nil
	}
	if startVal+lVal > srLen {
		lVal = srLen - startVal
	}

	return interpreter.StringOf(string(sr[startVal : startVal+lVal])), nil
}

type startsWithFunction struct {
	interpreter.BaseFunction
}

func newStartsWithFunction() *startsWithFunction {
	return &startsWithFunction{
		BaseFunction: interpreter.NewBaseFunction("startsWith", -1, 1, 1),
	}
}

func (f *startsWithFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return interpreter.BooleanOf(strings.HasPrefix(s.String(), ss.String())), nil
}

type endsWithFunction struct {
	interpreter.BaseFunction
}

func newEndsWithFunction() *endsWithFunction {
	return &endsWithFunction{
		BaseFunction: interpreter.NewBaseFunction("endsWith", -1, 1, 1),
	}
}

func (f *endsWithFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return interpreter.BooleanOf(strings.HasSuffix(s.String(), ss.String())), nil
}

type containsFunction struct {
	interpreter.BaseFunction
}

func newContainsFunction() *containsFunction {
	return &containsFunction{
		BaseFunction: interpreter.NewBaseFunction("contains", -1, 1, 1),
	}
}

func (f *containsFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return interpreter.BooleanOf(strings.Contains(s.String(), ss.String())), nil
}

type upperFunction struct {
	interpreter.BaseFunction
}

func newUpperFunction() *upperFunction {
	return &upperFunction{
		BaseFunction: interpreter.NewBaseFunction("upper", -1, 0, 0),
	}
}

func (f *upperFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return interpreter.NewString(strings.ToUpper(s.String())), nil
}

type lowerFunction struct {
	interpreter.BaseFunction
}

func newLowerFunction() *lowerFunction {
	return &lowerFunction{
		BaseFunction: interpreter.NewBaseFunction("lower", -1, 0, 0),
	}
}

func (f *lowerFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return interpreter.NewString(strings.ToLower(s.String())), nil
}

type replaceFunction struct {
	interpreter.BaseFunction
}

func newReplaceFunction() *replaceFunction {
	return &replaceFunction{
		BaseFunction: interpreter.NewBaseFunction("replace", -1, 2, 2),
	}
}

func (f *replaceFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	pattern, err := stringNode(args[0])
	if pattern == nil || err != nil {
		return nil, err
	}

	substitution, err := stringNode(args[1])
	if substitution == nil || err != nil {
		return nil, err
	}

	res := strings.ReplaceAll(s.String(), pattern.String(), substitution.String())
	return interpreter.StringOf(res), nil
}

type matchesFunction struct {
	interpreter.BaseFunction
}

func newMatchesFunction() *matchesFunction {
	return &matchesFunction{
		BaseFunction: interpreter.NewBaseFunction("matches", -1, 1, 1),
	}
}

func (f *matchesFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	re, err := stringNode(args[0])
	if re == nil || err != nil {
		return nil, err
	}

	b, err := regexp.MatchString(re.String(), s.String())
	if err != nil {
		return nil, err
	}
	return interpreter.BooleanOf(b), nil
}

type replaceMatchesFunction struct {
	interpreter.BaseFunction
}

func newReplaceMatchesFunction() *replaceMatchesFunction {
	return &replaceMatchesFunction{
		BaseFunction: interpreter.NewBaseFunction("replaceMatches", -1, 2, 2),
	}
}

func (f *replaceMatchesFunction) Execute(_ interpreter.ContextAccessor, node interface{}, args []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	regex, err := stringNode(args[0])
	if regex == nil || err != nil {
		return nil, err
	}

	substitution, err := stringNode(args[1])
	if substitution == nil || err != nil {
		return nil, err
	}

	re, err := regexp.Compile(regex.String())
	if err != nil {
		return nil, err
	}

	return interpreter.StringOf(re.ReplaceAllString(s.String(), substitution.String())), nil
}

type lengthFunction struct {
	interpreter.BaseFunction
}

func newLengthFunction() *lengthFunction {
	return &lengthFunction{
		BaseFunction: interpreter.NewBaseFunction("length", -1, 0, 0),
	}
}

func (f *lengthFunction) Execute(_ interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return interpreter.NewInteger(int32(utf8.RuneCountInString(s.String()))), nil
}

type toCharsFunction struct {
	interpreter.BaseFunction
}

func newToCharsFunction() *toCharsFunction {
	return &toCharsFunction{
		BaseFunction: interpreter.NewBaseFunction("toChars", -1, 0, 0),
	}
}

func (f *toCharsFunction) Execute(ctx interpreter.ContextAccessor, node interface{}, _ []interface{}, _ interpreter.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	if s.Length() == 0 {
		return interpreter.NewEmptyCol(), nil
	}

	col := ctx.NewCol()
	sr := []rune(s.String())
	for _, c := range sr {
		col.Add(interpreter.StringOf(string(c)))
	}

	return col, nil
}

func stringNode(node interface{}) (interpreter.StringAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if s, ok := value.(interpreter.StringAccessor); !ok {
		return nil, fmt.Errorf("expected string, but found %T", value)
	} else {
		return s, nil
	}
}
