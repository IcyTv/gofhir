package r4_test

import (
	"testing"

	"github.com/gofhir/fhirpath"
	"github.com/gofhir/fhirpath/r4"
	"github.com/stretchr/testify/assert"
)

// var subject = r4.JSONNode{
// 	"resourceType": "Observation",
// 	"valueDecimal": 123.45,
// 	"value":        "TODO remove",
// }

// func TestBasic(t *testing.T) {
// 	ctx := r4.NewR4Context(&r4.R4Model{})
// 	res, err := fhirpath.Execute(ctx, "value", subject)
// 	assert.Nil(t, err, "Error executing FHIRPath")
// 	assert.NotNil(t, res, "Result is nil")
// }

// func TestBasicWithRootType(t *testing.T) {
// 	ctx := r4.NewR4Context(&r4.R4Model{})
// 	res, err := fhirpath.Execute(ctx, "Observation.value", subject)
// 	assert.Nil(t, err, "Error executing FHIRPath")
// 	assert.NotNil(t, res, "Result is nil")
// }

func TestComplex(t *testing.T) {
	subject := map[string]interface{}{
		"resourceType": "Observation",
		"valueString":  "Hello, World!",
	}

	ctx := r4.NewR4Context(&r4.R4Model{})
	s := r4.NewTestNode("", subject)
	res, err := fhirpath.Execute(ctx, "Observation.value", s)
	assert.Nil(t, err, "Error executing FHIRPath")
	assert.NotNil(t, res, "Result is nil")
}
