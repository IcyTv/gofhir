package expression

import "github.com/gofhir/fhirpath/interpreter"

const delimitedIdentifierChar = '`'

func ExtractIdentifier(value string) string {
	resultingValue := value
	if len(resultingValue) > 1 && value[0] == delimitedIdentifierChar {
		resultingValue = resultingValue[1 : len(resultingValue)-1]
	}
	return resultingValue
}

func uniteCollections(ctx interpreter.ContextAccessor, n1 interface{}, n2 interface{}) interpreter.ColModifier {
	if n1 == nil && n2 == nil {
		return nil
	}

	c := ctx.NewCol()
	addUniqueCollectionItems(c, n1)
	addUniqueCollectionItems(c, n2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addUniqueCollectionItems(collection interpreter.ColModifier, node interface{}) {
	if node == nil {
		return
	}

	if c, ok := node.(interpreter.ColAccessor); ok {
		collection.AddAllUnique(c)
	} else {
		collection.AddUnique(node)
	}
}

func unwrapCollection(node interface{}) interface{} {
	if node == nil {
		return nil
	}

	if c, ok := node.(interpreter.ColAccessor); !ok {
		return node
	} else {
		count := c.Count()
		if count == 0 {
			return nil
		}
		if count == 1 {
			return c.Get(0)
		}
		return c
	}
}

func wrapCollection(ctx interpreter.ContextAccessor, node interface{}) interpreter.ColAccessor {
	if node == nil {
		return interpreter.EmptyColl
	}

	if col, ok := node.(interpreter.ColAccessor); ok {
		return col
	} else {
		return ctx.NewColWithItem(node)
	}
}

func emptyCollection(node interface{}) bool {
	if node == nil {
		return true
	}

	if col, ok := node.(interpreter.ColAccessor); ok {
		return col.Empty()
	}

	return false
}

func combineCollections(ctx interpreter.ContextAccessor, n1 interface{}, n2 interface{}) interpreter.ColModifier {
	if n1 == nil && n2 == nil {
		return nil
	}

	c := ctx.NewCol()
	addCollectionItems(c, n1)
	addCollectionItems(c, n2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addCollectionItems(collection interpreter.ColModifier, node interface{}) {
	if node == nil {
		return
	}
	if c, ok := node.(interpreter.ColAccessor); ok {
		collection.AddAll(c)
	} else {
		collection.Add(node)
	}
}
