package errors

import (
	"fmt"
	"log"
	"strings"
)

type Error struct {
	msg   string
	items []*ErrorItem
}

func (e *Error) String() string {
	log.Println("Error String method called")
	var b strings.Builder
	b.WriteString(e.msg)
	for _, item := range e.items {
		b.WriteString("\n- ")
		b.WriteString(item.String())
	}

	return b.String()
}

type ErrorItem struct {
	line   int
	column int
	msg    string
}

func NewError(msg string, item []*ErrorItem) *Error {
	return &Error{msg, item}
}

func NewErrorItem(line int, column int, msg string) *ErrorItem {
	return &ErrorItem{line, column, msg}
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Items() []*ErrorItem {
	return e.items
}

func (e *ErrorItem) Line() int {
	return e.line
}

func (e *ErrorItem) Column() int {
	return e.column
}

func (e *ErrorItem) Msg() string {
	return e.msg
}

func (e *ErrorItem) String() string {
	return fmt.Sprintf("%d:%d %s", e.line, e.column, e.msg)
}
