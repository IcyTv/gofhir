package r4

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gofhir/fhirpath/interpreter"
)

//go:generate go run generator/main.go

type R4Context struct {
	modelAdapter interpreter.ModelAdapter
	tracer       interpreter.Tracer
	node         interface{}
}

func NewR4Context(model *R4Model) interpreter.ContextAccessor {
	return &R4Context{modelAdapter: model}
}

func (t *R4Context) EnvVar(name string) (interface{}, bool) {
	if name == "ucum" {
		return interpreter.UCUMSystemURI, true
	}
	return nil, false
}

func (t *R4Context) ModelAdapter() interpreter.ModelAdapter {
	return t.modelAdapter
}

func (t *R4Context) NewCol() interpreter.ColModifier {
	return interpreter.NewCol(t.modelAdapter)
}

func (t *R4Context) NewColWithItem(item interface{}) interpreter.ColModifier {
	return interpreter.NewColWithItem(t.modelAdapter, item)
}

func (t *R4Context) ContextNode() interface{} {
	return t.node
}

func (t *R4Context) Tracer() interpreter.Tracer {
	return t.tracer
}

func (t *R4Model) Navigate(node interface{}, name string) (interface{}, error) {
	log.Println("Navigating", name, reflect.TypeOf(node))
	if n, ok := node.(R4ModelNodeInterface); ok {
		// if n.GetResourceType() == name {
		// 	return n, nil
		// } else {
		// 	return n.GetChild(name)
		// }
		return n.GetChild(name)
	} else {
		log.Printf("TODO: R4Model.Navigate: %T", node)
		return nil, fmt.Errorf("Cannot navigate to '%s' on %T", name, node)
	}
}

type R4ModelNodeInterface interface {
	GetResourceType() string
	GetChild(name string) (interface{}, error)
}

type TestNode struct {
	m    map[string]interface{}
	path string
}

func NewTestNode(path string, m map[string]interface{}) *TestNode {
	return &TestNode{m: m, path: path}
}

func (t *TestNode) GetResourceType() string {
	return t.m["resourceType"].(string)
}

func (t *TestNode) GetChild(name string) (interface{}, error) {
	var realPathName string
	if t.path == "" {
		realPathName = name
	} else {
		realPathName = t.path + "." + name
	}

	if t.m["resourceType"] != nil && name == t.m["resourceType"].(string) {
		return NewTestNode(realPathName, t.m), nil
	}

	if typeAliases[realPathName] != nil {
		for _, alias := range typeAliases[realPathName] {
			if val, ok := t.m[name+alias]; ok {
				// return NewTestNode(realPathName, val.(map[string]interface{})), nil
				asMap, ok := val.(map[string]interface{})
				if !ok {
					return val, nil
				} else {
					return NewTestNode(realPathName, asMap), nil
				}
			}
		}
	}

	if val, ok := t.m[name]; ok {
		// return NewTestNode(realPathName, val.(map[string]interface{})), nil
		asMap, ok := val.(map[string]interface{})
		if !ok {
			return val, nil
		} else {
			return NewTestNode(realPathName, asMap), nil
		}
	} else {
		return nil, fmt.Errorf("No such property '%s' on '%T'", name, t)
	}
}
