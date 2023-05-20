package visitors

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ErrorListener struct {
	antlr.DefaultErrorListener
	errorItemCollection *ErrorItemCollection
}

func NewErrorListener(errorItemCollection *ErrorItemCollection) *ErrorListener {
	l := new(ErrorListener)
	l.errorItemCollection = errorItemCollection
	return l
}

func (l *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errorItemCollection.AddError(line, column, msg)
}
