package visitors

import (
	"github.com/gofhir/fhirpath/errors"
)

type ErrorItemCollection struct {
	items []*errors.ErrorItem
}

func NewErrorItemCollection() *ErrorItemCollection {
	items := make([]*errors.ErrorItem, 0)
	return &ErrorItemCollection{items}
}

func (c *ErrorItemCollection) AddError(line int, column int, msg string) {
	item := errors.NewErrorItem(line, column, msg)
	c.items = append(c.items, item)
}

func (c *ErrorItemCollection) HasErrors() bool {
	return len(c.items) > 0
}

func (c *ErrorItemCollection) Items() []*errors.ErrorItem {
	return c.items
}
