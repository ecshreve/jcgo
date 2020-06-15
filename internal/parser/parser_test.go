package parser_test

import (
	"errors"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/internal/object"
	"github.com/ecshreve/jcgo/internal/parser"
	"github.com/ecshreve/jcgo/internal/testdata"
)

func TestValidateConfig(t *testing.T) {
	testcases := []struct {
		description string
		input       *parser.Config
		valid       bool
	}{
		{
			description: "no input file",
			input: &parser.Config{
				Infile: "",
			},
			valid: false,
		},
		{
			description: "non-json input file",
			input: &parser.Config{
				Infile: "testfilename.csv",
			},
			valid: false,
		},
		{
			description: "valid input file",
			input: &parser.Config{
				Infile: "testfilename.json",
			},
			valid: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			err := testcase.input.Validate()
			if !testcase.valid {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	testcases := []struct {
		description    string
		inputArgs      []string
		expectedConfig *parser.Config
		expectedOutput string
		expectError    bool
	}{
		{
			description: "basic valid args",
			inputArgs:   []string{"dummy", "--infile=testfile.json"},
			expectedConfig: &parser.Config{
				Infile:  "testfile.json",
				Outfile: "testfile.output.csv",
				Args:    []string{},
			},
			expectedOutput: "",
		},
		{
			description:    "invalid command line arg",
			inputArgs:      []string{"dummy", "--somearg=someval"},
			expectedConfig: nil,
			expectedOutput: "flag provided but not defined: -somearg",
			expectError:    true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actualConfig, actualOutput, err := parser.ParseArgs(testcase.inputArgs)
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expectedConfig, actualConfig)

			actualOutputPrefix := strings.Split(actualOutput, "\n")[0]
			assert.Equal(t, testcase.expectedOutput, actualOutputPrefix)
		})
	}
}

func TestHandleParseError(t *testing.T) {
	testcases := []struct {
		description      string
		inputOutput      string
		inputError       error
		expectedExitCode int
		expectError      bool
	}{
		{
			description:      "valid input",
			inputOutput:      "dummy string",
			inputError:       nil,
			expectedExitCode: 0,
			expectError:      false,
		},
		{
			description:      "help error",
			inputOutput:      "dummy string",
			inputError:       flag.ErrHelp,
			expectedExitCode: 2,
			expectError:      false,
		},
		{
			description:      "generic error",
			inputOutput:      "dummy string",
			inputError:       errors.New("generic error"),
			expectedExitCode: 1,
			expectError:      true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actualExitCode, err := parser.HandleParseError(testcase.inputOutput, testcase.inputError)
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expectedExitCode, actualExitCode)
		})
	}
}

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
	testcases := []struct {
		description         string
		input               *parser.Config
		expectedOutfilePath string
		expectError         bool
	}{
		{
			description: "valid infile, default outfile",
			input: &parser.Config{
				Infile:  "../testdata/jsontest.json",
				Outfile: "../testdata/jsontest.output.csv",
			},
			expectedOutfilePath: "../testdata/jsontest.output.csv",
		},
		{
			description: "valid infile, explicit outfile",
			input: &parser.Config{
				Infile:  "../testdata/jsontest.json",
				Outfile: "../testdata/explicit_output_file.csv",
			},
			expectedOutfilePath: "../testdata/explicit_output_file.csv",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actualFile, err := parser.Convert(testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expectedOutfilePath, actualFile.Name())
		})
	}
}

func TestSetDefaultOutfile(t *testing.T) {
	testcases := []struct {
		description string
		input       *parser.Config
		expected    *parser.Config
		expectError bool
	}{
		{
			description: "outfile already exists, expect error",
			input: &parser.Config{
				Infile:  "testinfile.json",
				Outfile: "testoutfile.csv",
			},
			expected: &parser.Config{
				Infile:  "testinfile.json",
				Outfile: "testoutfile.csv",
			},
			expectError: true,
		},
		{
			description: "outfile doesn't exist, expect default outfile and no error",
			input: &parser.Config{
				Infile: "testinfile.json",
			},
			expected: &parser.Config{
				Infile:  "testinfile.json",
				Outfile: "testinfile.output.csv",
			},
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			err := testcase.input.SetDefaultOutfile()
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expected, testcase.input)
		})
	}
}
