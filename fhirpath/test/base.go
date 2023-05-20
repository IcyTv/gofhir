package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gofhir/fhirpath/interpreter"
)

var testBaseTypeSpec = interpreter.NewTypeSpec(interpreter.NewFQTypeName("base", "TEST"))
var testTypeSpec = interpreter.NewTypeSpecWithBase(interpreter.NewFQTypeName("type1", "TEST"), testBaseTypeSpec)
var testElementTypeSpec = interpreter.NewTypeSpecWithBase(interpreter.NewFQTypeName("Element", "TEST"), testBaseTypeSpec)

type testModelNode struct {
	value float64
	sys   bool
}

type testModelErrorNode struct {
}

type testModelNodeAccessor interface {
	testValue() float64
	testSys() bool
}

type testModelErrorNodeAccessor interface {
	testErr() bool
}

func NewTestModelNode(value float64, sys bool) *testModelNode {
	return &testModelNode{value, sys}
}

func NewTestModelErrorNode() *testModelErrorNode {
	return &testModelErrorNode{}
}

func (n *testModelNode) testValue() float64 {
	return n.value
}

func (n *testModelNode) testSys() bool {
	return n.sys
}

func (n *testModelErrorNode) testErr() bool {
	return true
}

type testModel struct {
	t *testing.T
}

func newTestModel(t *testing.T) interpreter.ModelAdapter {
	return &testModel{t}
}

func (a *testModel) CastToSystem(node interface{}) (interpreter.AnyAccessor, error) {
	if _, ok := node.(map[string]interface{}); ok {
		return nil, nil
	}

	if n, ok := node.(testModelNodeAccessor); !ok {
		if _, ok = node.(testModelErrorNodeAccessor); ok {
			return nil, fmt.Errorf("error node cannot be converted")
		}
		a.t.Errorf("not a test model node: %T", node)
		return nil, nil
	} else {
		if n.testSys() {
			return interpreter.NewDecimalFloat64(n.testValue()), nil
		}
		return nil, nil
	}
}

func (a *testModel) AsType(node interface{}, name interpreter.FQTypeNameAccessor) (interface{}, error) {
	if n, ok := node.(testModelNodeAccessor); !ok {
		a.t.Errorf("not a test model node: %T", node)
		return nil, nil
	} else {
		if name.Namespace() == "Other" {
			return nil, fmt.Errorf("unsupported conversion")
		}
		if name.Name() == "decimal" && name.Namespace() == "Test" {
			return n, nil
		}
		if name.Name() == "Number" && name.Namespace() == "System" {
			return interpreter.NewDecimalFloat64(n.testValue()), nil
		}
		return nil, nil
	}
}

func (a *testModel) TypeSpec(node interface{}) interpreter.TypeSpecAccessor {
	if _, ok := node.(map[string]interface{}); ok {
		return testElementTypeSpec
	}

	if n, ok := node.(testModelNodeAccessor); !ok {
		if _, ok := node.(testModelErrorNodeAccessor); ok {
			return interpreter.NewTypeSpec(interpreter.NewFQTypeName("Error", "Test"))
		}
		if _, ok := node.(interpreter.StringAccessor); !ok {
			a.t.Errorf("not a test model node: %T", node)
		}
		return interpreter.UndefinedTypeSpec
	} else {
		if n.testSys() {
			a.t.Errorf("type of system node must not be requested")
		}
		return testTypeSpec
	}
}

func (a *testModel) Equal(node1 interface{}, node2 interface{}) bool {
	if m1, ok := node1.(map[string]interface{}); ok {
		if m2, ok := node2.(map[string]interface{}); !ok {
			return false
		} else {
			id := m1["id"]
			return id != nil && id == m2["id"]
		}
	}

	n1, ok := node1.(testModelNodeAccessor)
	if !ok {
		return false
	}
	n2, ok := node2.(testModelNodeAccessor)
	if !ok {
		return false
	}

	if n1.testSys() || n2.testSys() {
		a.t.Errorf("equality of system node must not be requested")
	}
	return n1.testValue() == n2.testValue()
}

func (a *testModel) Equivalent(node1 interface{}, node2 interface{}) bool {
	n1, ok := node1.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node1)
	}
	n2, ok := node2.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node2)
	}

	if n1.testSys() || n2.testSys() {
		a.t.Errorf("equality of system node must not be requested")
	}
	return int64(n1.testValue()) == int64(n2.testValue())
}

func (a *testModel) Navigate(node interface{}, name string) (interface{}, error) {
	var result interface{}
	model, ok := node.(map[string]interface{})
	if ok {
		var found bool
		// a.t.Fatal(fmt.Sprintf("cannot be cast to map: %T", node))
		result, found = model[name]
		if !found {
			return nil, fmt.Errorf("path cannot be evaluated on model: %s", name)
		}
	}

	if n, ok := node.(interpreter.QuantityAccessor); ok {
		if name == "value" {
			return n.Value(), nil
		}
		if name == "unit" {
			return n.Unit(), nil
		}
		return nil, fmt.Errorf("path cannot be evaluated on quantity: %s", name)
	}

	return result, nil
}

func (a *testModel) Children(node interface{}) (interpreter.ColAccessor, error) {
	if _, ok := node.(interpreter.AnyAccessor); ok {
		return nil, nil
	}

	model, ok := node.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot be cast to map: %T", node)
	}

	res := interpreter.NewCol(a)
	keys := make([]string, 0)
	for k := range model {
		if k == "errorCollection" {
			return NewErrorCollection(), nil
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := model[k]
		if v != nil {
			res.Add(model[k])
		}
	}
	return res, nil
}

type testContext struct {
	modelAdapter interpreter.ModelAdapter
	tracer       interpreter.Tracer
	node         interface{}
}

func NewTestContext(t *testing.T) interpreter.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t)}
}

func NewTestContextWithNode(t *testing.T, node interface{}) interpreter.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t), node: node}
}

func NewTestContextWithNodeAndTracer(t *testing.T, node interface{}, tracer interpreter.Tracer) interpreter.ContextAccessor {
	return &testContext{
		modelAdapter: newTestModel(t),
		tracer:       tracer,
		node:         node,
	}
}

func (t *testContext) EnvVar(name string) (interface{}, bool) {
	if name == "ucum" {
		return interpreter.UCUMSystemURI, true
	}
	return nil, false
}

func (t *testContext) ModelAdapter() interpreter.ModelAdapter {
	return t.modelAdapter
}

func (t *testContext) NewCol() interpreter.ColModifier {
	return interpreter.NewCol(t.modelAdapter)
}

func (t *testContext) NewColWithItem(item interface{}) interpreter.ColModifier {
	return interpreter.NewColWithItem(t.modelAdapter, item)
}

func (t *testContext) ContextNode() interface{} {
	return t.node
}

func (t *testContext) Tracer() interpreter.Tracer {
	return t.tracer
}

type errorCollection struct {
}

func NewErrorCollection() interpreter.ColAccessor {
	return &errorCollection{}
}

func (c *errorCollection) DataType() interpreter.DataTypes {
	return interpreter.ColDataType
}

func (c *errorCollection) Empty() bool {
	return false
}

func (c *errorCollection) Count() int {
	return 3
}

func (c *errorCollection) Get(i int) interface{} {
	switch i {
	case 0:
		return interpreter.NewString("test item 1")
	case 1:
		return NewTestModelErrorNode()
	case 2:
		return interpreter.NewString("test item 2")
	default:
		panic("invalid item index")
	}
}

func (c *errorCollection) Contains(node interface{}) bool {
	if _, ok := node.(testModelErrorNodeAccessor); ok {
		return true
	}
	return false
}

func (c *errorCollection) TypeSpec() interpreter.TypeSpecAccessor {
	panic("implement me")
}

func (c *errorCollection) TypeInfo() interpreter.TypeInfoAccessor {
	panic("implement me")
}

func (c *errorCollection) Source() interface{} {
	panic("implement me")
}

func (c *errorCollection) Equal(_ interface{}) bool {
	panic("implement me")
}

func (c *errorCollection) Equivalent(_ interface{}) bool {
	panic("implement me")
}

func (c *errorCollection) ItemTypeSpec() interpreter.TypeSpecAccessor {
	panic("implement me")
}
