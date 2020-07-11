package parser_test

import (
	"errors"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/internal/parser"
)

func TestValidateConfig(t *testing.T) {
	testcases := []struct {
		description string
		input       *parser.Config
		valid       bool
	}{
		{
			description: "no infile",
			input: &parser.Config{
				Infile: "",
			},
			valid: false,
		},
		{
			description: "non-json infile",
			input: &parser.Config{
				Infile: "testfilename.csv",
			},
			valid: false,
		},
		{
			description: "valid infile, valid outfile",
			input: &parser.Config{
				Infile:  "testfilename.json",
				Outfile: "testfilename.output.csv",
			},
			valid: true,
		},
		{
			description: "valid infile, no outfile",
			input: &parser.Config{
				Infile: "testfilename.json",
			},
			valid: false,
		},
		{
			description: "valid infile, invalid outfile",
			input: &parser.Config{
				Infile:  "testfilename.json",
				Outfile: "badoutfile.jpeg",
			},
			valid: false,
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

func TestFileOperations(t *testing.T) {
	// Verify we can read a valid JSON file.
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
		{"one", "two"},
		{"one_one", "two_two"},
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
