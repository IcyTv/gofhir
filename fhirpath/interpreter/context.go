package interpreter

type ModelAdapter interface {
	AsType(node interface{}, name FQTypeNameAccessor) (interface{}, error)
	// CastToSystem(node interface{}) (AnyAccessor, error)
	TypeSpec(node interface{}) TypeSpecAccessor
	Equal(node1 interface{}, node2 interface{}) bool
	Equivalent(node1 interface{}, node2 interface{}) bool
	Navigate(node interface{}, name string) (interface{}, error)
	Children(node interface{}) (ColAccessor, error)
}

func ModelTypeSpec(adapter ModelAdapter, node interface{}) TypeSpecAccessor {
	if node == nil {
		return nil
	}

	if n, ok := node.(AnyAccessor); ok {
		return n.TypeSpec()
	}
	return adapter.TypeSpec(node)
}

func HasModelType(adapter ModelAdapter, node interface{}, name FQTypeNameAccessor) bool {
	if node == nil {
		return false
	}

	namespace := name.Namespace()
	if systemNamespace(namespace) {
		if n, ok := node.(AnyAccessor); ok {
			if n.TypeSpec().ExtendsName(name) {
				return true
			}
		}
	}

	if namespace != NamespaceName {
		if adapter.TypeSpec(node).ExtendsName(name) {
			return true
		}
	}

	return false
}

func CastModelType(adapter ModelAdapter, node interface{}, name FQTypeNameAccessor) (interface{}, error) {
	if node == nil {
		return node, nil
	}

	sysNode, sys := node.(AnyAccessor)
	if systemNamespace(name.Namespace()) && sys && sysNode.TypeSpec().ExtendsName(name) {
		return node, nil
	}
	if sys {
		// system node cannot be casted by model adapter
		return nil, nil
	}

	return adapter.AsType(node, name)
}

func ModelEqual(adapter ModelAdapter, node1 interface{}, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode1, _ := node1.(AnyAccessor)
	if sysNode1 != nil {
		sysNode2, _ := node2.(AnyAccessor)
		if sysNode2 != nil {
			return sysNode1.Equal(sysNode2)
		}
	}

	return adapter != nil && adapter.Equal(node1, node2)
}

func ModelEquivalent(adapter ModelAdapter, node1 interface{}, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode1, _ := node1.(AnyAccessor)
	if sysNode1 != nil {
		sysNode2, _ := node2.(AnyAccessor)
		if sysNode2 != nil {
			return sysNode1.Equivalent(sysNode2)
		}
	}

	return adapter != nil && adapter.Equivalent(node1, node2)
}

func SystemAnyTypeEqual(node1 AnyAccessor, node2 interface{}) bool {
	if node1 == nil || node2 == nil {
		return false
	}
	if sysNode2, ok := node2.(AnyAccessor); ok {
		return node1.TypeSpec().EqualType(sysNode2.TypeSpec())
	}
	return false
}

func SystemAnyEqual(node1 AnyAccessor, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode2, _ := node2.(AnyAccessor)
	if sysNode2 != nil {
		return false
	}

	return Equal(node1, sysNode2)
}

type Tracer interface {
	Enabled(name string) bool
	Trace(name string, col ColAccessor)
}

type ContextAccessor interface {
	EnvVar(name string) (interface{}, bool)
	ContextNode() interface{}
	ModelAdapter() ModelAdapter
	NewCol() ColModifier
	NewColWithItem(item interface{}) ColModifier
	Tracer() Tracer
}

func systemNamespace(name string) bool {
	return len(name) == 0 || name == NamespaceName
}
