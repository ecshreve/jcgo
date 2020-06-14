package parser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/internal/object"
	"github.com/ecshreve/jcgo/internal/parser"
	"github.com/ecshreve/jcgo/internal/testdata"
)

func TestTransform(t *testing.T) {
	data := testdata.NewObjectTestData()

	testcases := []struct {
		description string
		input       map[string]interface{}
		expected    object.Object
		expectError bool
	}{
		{
			description: "invalid input",
			input: map[string]interface{}{
				"key1": int64(1),
				"key2": "val2",
				"key3": "val3",
			},
			expected:    nil,
			expectError: true,
		},
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
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestFileOperations(t *testing.T) {
	// Verify we can read a valid json file.
	raw, err := parser.ReadJSONFile("../testdata/jsontest.json")
	assert.NoError(t, err)
	assert.NotNil(t, raw)

	// Expect error if file doesn't exist.
	raw, err = parser.ReadJSONFile("../testdata/nonexistentfile.json")
	assert.Error(t, err)
	assert.Nil(t, raw)

	// Expect error if file is malformed.
	raw, err = parser.ReadJSONFile("../testdata/jsontest_bad.json")
	assert.Error(t, err)
	assert.Nil(t, raw)

	// Verify we can write valid data to CSV.
	data := [][]string{
		[]string{"one", "two"},
		[]string{"one_one", "two_two"},
	}
	file, err := parser.WriteCSVFile(data, "../testdata/testoutput.csv")
	assert.NoError(t, err)
	assert.NotNil(t, file)

	// Expect error for a bad path.
	file, err = parser.WriteCSVFile(data, "../testdata/nonexistentdirectory/testoutput.xls")
	assert.Error(t, err)
	assert.Nil(t, file)
	os.Remove("../testdata/testoutput.csv")
}

func TestConvert(t *testing.T) {
	file, err := parser.Convert("../testdata/jsontest.json")
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, "../testdata/jsontest.output.csv", file.Name())
}
