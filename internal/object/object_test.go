package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	oo "github.com/ecshreve/jcgo/internal/object"
	"github.com/ecshreve/jcgo/internal/testdata"
)

func TestObjectCreation(t *testing.T) {
	testcases := []struct {
		description string
		input       interface{}
		expectError bool
	}{
		{
			description: "create nil StringObj",
			input:       nil,
			expectError: false,
		},
		{
			description: "create StringObj",
			input:       "this is a string",
			expectError: false,
		},
		{
			description: "create BoolObj",
			input:       true,
			expectError: false,
		},
		{
			description: "create FloatObj",
			input:       float64(10.0),
			expectError: false,
		},
		{
			description: "invalid scalar input",
			input:       int64(1),
			expectError: true,
		},
		{
			description: "create MapObj",
			input:       map[string]interface{}{"key": "value"},
			expectError: false,
		},
		{
			description: "invalid type in map",
			input:       map[string]interface{}{"key": int64(1)},
			expectError: true,
		},
		{
			description: "create SliceObj",
			input:       []interface{}{"str1", "str2"},
			expectError: false,
		},
		{
			description: "invalid type in slice",
			input:       []interface{}{"str1", int64(1)},
			expectError: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			_, err := oo.FromInterface("", testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
		})
	}
}

func TestParseObject(t *testing.T) {
	data := testdata.NewObjectTestData()

	testcases := []struct {
		description string
		input       oo.Object
		expected    [][]string
	}{
		{
			description: "simple string",
			input: &oo.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			expected: [][]string{
				{"key1"},
				{"val1"},
			},
		},
		{
			description: "simple map",
			input:       data.SimpleMapObj,
			expected: [][]string{
				{"key1", "key2", "key3"},
				{"val1", "val2", "val3"},
			},
		},
		{
			description: "simple map with all possible data types",
			input:       data.SimpleAllTypesMapObj,
			expected: [][]string{
				{"key1", "key2", "key3", "key4"},
				{"val1", "5", "5.5", "true"},
			},
		},
		{
			description: "simple nested map",
			input:       data.SimpleNestedMapObj,
			expected: [][]string{
				{"outer1_inner1", "outer1_inner2", "outer2"},
				{"innerval1", "innerval2", "outerval2"},
			},
		},
		{
			description: "double nested map",
			input:       data.DoubleNestedMapObj,
			expected: [][]string{
				{"outer1_inner1", "outer1_nestedmap_nested1", "outer1_nestedmap_nested2", "outer2"},
				{"innerval1", "nestedval1", "nestedval2", "outerval2"},
			},
		},
		{
			description: "map with simple slice",
			input:       data.SimpleSliceMapObj,
			expected: [][]string{
				{"data_key1", "data_key2", "data_key3"},
				{"val1", "val2", "val3"},
				{"val4", "val5", "val6"},
			},
		},
		{
			description: "map with slice with nested map",
			input:       data.SimpleSliceNestedMapObj,
			expected: [][]string{
				{"data_key1", "data_nestedmap_nested1", "data_nestedmap_nested2"},
				{"val1", "nestedval1", "nestedval2"},
				{"val4", "nestedval3", "nestedval4"},
			},
		},
		{
			description: "complex map with nested maps and slices",
			input:       data.ComplexMapObj,
			expected: [][]string{
				{"data_key1", "data_nestedmap_nested1", "data_nestedmap_nested2", "data_nestedslice_after_nestedslicemap1", "data_nestedslice_after_nestedslicemap2", "data_nestedslice_before_nestedslicemap1", "data_nestedslice_before_nestedslicemap2"},
				{"val1", "nestedval1", "nestedval2", "nestedslicemapval5", "nestedslicemapval6", "nestedslicemapval1", "nestedslicemapval2"},
				{"val1", "nestedval1", "nestedval2", "nestedslicemapval7", "nestedslicemapval8", "nestedslicemapval3", "nestedslicemapval4"},
				{"val1", "nestedval1", "nestedval2", "nestedslicemapval15", "nestedslicemapval16", "nestedslicemapval11", "nestedslicemapval12"},
				{"val1", "nestedval1", "nestedval2", "nestedslicemapval17", "nestedslicemapval18", "nestedslicemapval13", "nestedslicemapval14"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual, _ := testcase.input.Parse()
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
