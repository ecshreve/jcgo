package parser_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/parser"
	"github.com/ecshreve/jcgo/parser/testdata"
)

func TestObjectCreation(t *testing.T) {
	testcases := []struct {
		description string
		input       interface{}
		expectError bool
	}{
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
			_, err := parser.ObjectFromInterface("", testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
		})
	}
}

func TestParseObject(t *testing.T) {
	data := testdata.NewObjectTestData()
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.Indent = "  "

	testcases := []struct {
		description string
		input       parser.Object
		expected    [][]string
	}{
		{
			description: "simple string",
			input: &parser.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			expected: [][]string{
				[]string{"key1"},
				[]string{"val1"},
			},
		},
		{
			description: "simple map",
			input:       data.SimpleMapObj,
			expected: [][]string{
				[]string{"key1", "key2", "key3"},
				[]string{"val1", "val2", "val3"},
			},
		},
		{
			description: "simple map with all possible data types",
			input:       data.SimpleAllTypesMapObj,
			expected: [][]string{
				[]string{"key1", "key2", "key3", "key4"},
				[]string{"val1", "5", "5.5", "true"},
			},
		},
		{
			description: "simple nested map",
			input:       data.SimpleNestedMapObj,
			expected: [][]string{
				[]string{"outer1_inner1", "outer1_inner2", "outer2"},
				[]string{"innerval1", "innerval2", "outerval2"},
			},
		},
		{
			description: "double nested map",
			input:       data.DoubleNestedMapObj,
			expected: [][]string{
				[]string{"outer1_inner1", "outer1_nestedmap_nested1", "outer1_nestedmap_nested2", "outer2"},
				[]string{"innerval1", "nestedval1", "nestedval2", "outerval2"},
			},
		},
		{
			description: "map with simple slice",
			input:       data.SimpleSliceMapObj,
			expected: [][]string{
				[]string{"data_key1", "data_key2", "data_key3"},
				[]string{"val1", "val2", "val3"},
				[]string{"val4", "val5", "val6"},
			},
		},
		{
			description: "map with slice with nested map",
			input:       data.SimpleSliceNestedMapObj,
			expected: [][]string{
				[]string{"data_key1", "data_nestedmap_nested1", "data_nestedmap_nested2"},
				[]string{"val1", "nestedval1", "nestedval2"},
				[]string{"val4", "nestedval3", "nestedval4"},
			},
		},
		{
			description: "complex map with nested maps and slices",
			input:       data.ComplexMapObj,
			expected: [][]string{
				[]string{"data_key1", "data_nestedmap_nested1", "data_nestedmap_nested2", "data_nestedslice_after_nestedslicemap1", "data_nestedslice_after_nestedslicemap2", "data_nestedslice_before_nestedslicemap1", "data_nestedslice_before_nestedslicemap2"},
				[]string{"val1", "nestedval1", "nestedval2", "nestedslicemapval5", "nestedslicemapval6", "nestedslicemapval1", "nestedslicemapval2"},
				[]string{"val1", "nestedval1", "nestedval2", "nestedslicemapval7", "nestedslicemapval8", "nestedslicemapval3", "nestedslicemapval4"},
				[]string{"val1", "nestedval1", "nestedval2", "nestedslicemapval15", "nestedslicemapval16", "nestedslicemapval11", "nestedslicemapval12"},
				[]string{"val1", "nestedval1", "nestedval2", "nestedslicemapval17", "nestedslicemapval18", "nestedslicemapval13", "nestedslicemapval14"},
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
