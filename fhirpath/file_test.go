package fhirpath

import (
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/gofhir/fhirpath/interpreter"
	"github.com/gofhir/fhirpath/r4"
	"github.com/gofhir/fhirpath/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestFhirpathJSFiles(t *testing.T) {
	t.Parallel()

	dir, err := os.ReadDir("../assets/test/fhirpath/cases")

	assert.NoError(t, err, "Error reading file")

	for _, file := range dir {
		generateTestForEntry(file, t)
	}
}

func generateTestForEntry(entry fs.DirEntry, t *testing.T) {
	name := entry.Name()
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		bytes, err := os.ReadFile("../assets/test/fhirpath/cases/" + name)
		assert.NoError(t, err, "Error reading file %s", name)
		var out map[string]interface{}
		err = yaml.Unmarshal(bytes, &out)
		assert.NoError(t, err, "Error unmarshalling yaml file %s", name)

		var subject interface{}
		if subjectValue, ok := out["subject"]; ok {
			subject = subjectValue
		} else {
			subject = map[string]interface{}{}
		}

		tests := out["tests"].([]interface{})

		for _, test := range tests {
			testCaseOrGroup := test.(map[string]interface{})

			for key, _ := range testCaseOrGroup {
				if strings.HasPrefix(key, "group") {
					generateTestsForGroup(testCaseOrGroup, subject, t)
				} else {
					generateTest(testCaseOrGroup, subject, t)
				}
			}
		}
	})
}

func generateTestsForGroup(group map[string]interface{}, subject interface{}, t *testing.T) {
	for _, value := range group {
		// name := strings.TrimPrefix(key, "group:")
		for _, test := range value.([]interface{}) {
			generateTest(test.(map[string]interface{}), subject, t)
		}
	}
}

func generateTest(testCase map[string]interface{}, subject interface{}, t *testing.T) {
	expression := testCase["expression"].(string)

	var desc string
	if testCase["desc"] != nil {
		desc = testCase["desc"].(string)
	} else {
		desc = expression
	}

	isError := testCase["error"] != nil && testCase["error"].(bool)

	var result []interface{}
	if testCase["result"] != nil {
		result = testCase["result"].([]interface{})
	}

	var model string
	if testCase["model"] != nil {
		model = testCase["model"].(string)
	}

	t.Run(desc, func(t *testing.T) {
		var s interface{}
		var ctx interpreter.ContextAccessor
		switch model {
		case "r4":
			ctx = r4.NewR4Context(&r4.R4Model{})
			s = r4.NewTestNode("", subject.(map[string]interface{}))
		case "":
			ctx = test.NewTestContext(t)
			s = subject
		default:
			t.Logf("Unknown model %s", model)
			t.SkipNow()
			return
		}

		res, err := Execute(ctx, expression, s)
		if isError {
			assert.Error(t, err, "Error expected")
		} else {
			if !assert.Nil(t, err, "Error not expected") {
				return
			}

			if !assert.Equal(t, len(result), res.Count(), "Result count mismatch") {
				return
			}

			for i, r := range result {
				switch r.(type) {
				case bool:
					c := res.Get(i).(interpreter.BooleanAccessor).Bool()
					assert.Equal(t, r, c, "Boolean result mismatch")
				case float64:
					c := res.Get(i).(interpreter.NumberAccessor).Float64()
					assert.Equal(t, r, c, "float result mismatch")
				case int:
					c := int(res.Get(i).(interpreter.NumberAccessor).Int())
					assert.Equal(t, r, c, "Integer result mismatch")
				case string:
					c := res.Get(i).(interpreter.Stringifier).String()
					assert.Equal(t, r, c, "String result mismatch")
				default:
					t.Errorf("Unknown result type %T", r)
				}
			}
		}
		_ = res
	})
}
