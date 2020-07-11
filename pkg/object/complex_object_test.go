package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	oo "github.com/ecshreve/jcgo/pkg/object"
	"github.com/ecshreve/jcgo/pkg/testdata"
)

func TestComplexObject(t *testing.T) {
	data := testdata.NewObjectTestData()

	testcases := []struct {
		description string
		input       interface{}
		expected    [][]string
	}{
		{
			description: "parse simple array of maps",
			input:       data.SimpleArrayMapInput,
			expected: [][]string{
				[]string{"key1", "key2", "key3"},
				[]string{"val1", "val2", "val3"},
				[]string{"val4", "val5", "val6"},
				[]string{"val7", "val8", "val9"},
			},
		},
		{
			description: "parse simple map of arrays",
			input:       data.SimpleMapArrayInput,
			expected: [][]string{
				[]string{"outer1", "outer2"},
				[]string{"val1", "val4"},
				[]string{"val2", "val5"},
				[]string{"val3", "val6"},
			},
		},
		{
			description: "parse complex input",
			input:       data.ComplexInput1,
			expected: [][]string{
				{"outer1_key1", "outer1_key2", "outer1_key3", "outer2_key1", "outer2_key2", "outer2_key3"},
				{"val1", "val2", "nval1", "val3", "val4", "nval3"},
				{"val1", "val2", "nval2", "val3", "val4", "nval4"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			obj, err := oo.FromInterface("", testcase.input)
			assert.NoError(t, err)

			parsed, err := obj.Parse()
			assert.NoError(t, err)

			assert.Equal(t, testcase.expected, parsed)
		})
	}
}
