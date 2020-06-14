package parser_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/parser"
	"github.com/ecshreve/jcgo/parser/testdata"
)

func TestTransform(t *testing.T) {
	data := testdata.NewObjectTestData()
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.Indent = "  "

	testcases := []struct {
		description string
		input       map[string]interface{}
		expected    parser.Object
	}{
		{
			description: "simple map",
			input: map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
				"key3": "val3",
			},
			expected: data.SimpleMapObj,
		},
		{
			description: "simple map with all possible data types",
			input: map[string]interface{}{
				"key1": "val1",
				"key2": float64(5),
				"key3": float64(5.5),
				"key4": true,
			},
			expected: data.SimpleAllTypesMapObj,
		},
		{
			description: "simple nested map",
			input: map[string]interface{}{
				"outer1": map[string]interface{}{
					"inner1": "innerval1",
					"inner2": "innerval2",
				},
				"outer2": "outerval2",
			},
			expected: data.SimpleNestedMapObj,
		},
		{
			description: "double nested map",
			input: map[string]interface{}{
				"outer1": map[string]interface{}{
					"inner1": "innerval1",
					"nestedmap": map[string]interface{}{
						"nested1": "nestedval1",
						"nested2": "nestedval2",
					},
				},
				"outer2": "outerval2",
			},
			expected: data.DoubleNestedMapObj,
		},
		{
			description: "map with simple slice",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"key2": "val2",
						"key3": "val3",
					},
					map[string]interface{}{
						"key1": "val4",
						"key2": "val5",
						"key3": "val6",
					},
				},
			},
			expected: data.SimpleSliceMapObj,
		},
		{
			description: "map with slice with nested map",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval1",
							"nested2": "nestedval2",
						},
					},
					map[string]interface{}{
						"key1": "val4",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval3",
							"nested2": "nestedval4",
						},
					},
				},
			},
			expected: data.SimpleSliceNestedMapObj,
		},
		{
			description: "complex map with nested maps and slices",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval1",
							"nested2": "nestedval2",
						},
						"nestedslice": []interface{}{
							map[string]interface{}{
								"before": []interface{}{
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval1",
										"nestedslicemap2": "nestedslicemapval2",
									},
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval3",
										"nestedslicemap2": "nestedslicemapval4",
									},
								},
								"after": []interface{}{
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval5",
										"nestedslicemap2": "nestedslicemapval6",
									},
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval7",
										"nestedslicemap2": "nestedslicemapval8",
									},
								},
							},
							map[string]interface{}{
								"before": []interface{}{
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval11",
										"nestedslicemap2": "nestedslicemapval12",
									},
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval13",
										"nestedslicemap2": "nestedslicemapval14",
									},
								},
								"after": []interface{}{
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval15",
										"nestedslicemap2": "nestedslicemapval16",
									},
									map[string]interface{}{
										"nestedslicemap1": "nestedslicemapval17",
										"nestedslicemap2": "nestedslicemapval18",
									},
								},
							},
						},
					},
				},
			},
			expected: data.ComplexMapObj,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual, err := parser.Transform(testcase.input)
			assert.NoError(t, err)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestParse(t *testing.T) {
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
