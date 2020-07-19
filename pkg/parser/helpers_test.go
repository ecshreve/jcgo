package parser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func TestReadJSONFile(t *testing.T) {
	testcases := []struct {
		description string
		infilePath  string
		expectError bool
	}{
		{
			description: "expect error if infile isn't a json file",
			infilePath:  "../testdata/filename.somenonjsonextension",
			expectError: true,
		},
		{
			description: "expect error if infile doesn't exist",
			infilePath:  "../testdata/nonexistentfile.json",
			expectError: true,
		},
		{
			description: "expect error if infile is malformed",
			infilePath:  "../testdata/jsontest_bad.json",
			expectError: true,
		},
		{
			description: "expect success for valid infile",
			infilePath:  "../testdata/jsontest.json",
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			rawData, err := parser.ReadJSONFile(testcase.infilePath)
			if testcase.expectError {
				assert.Error(t, err)
				assert.Nil(t, rawData)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, rawData)
			}
		})
	}
}

func TestWriteJSONFile(t *testing.T) {
	// Verify we can write valid data to CSV.
	data := [][]string{
		{"one", "two"},
		{"one_one", "two_two"},
	}

	// Expect error for a bad path.
	badOutfilePath := "../testdata/nonexistentdirectory/testoutput.xls"
	outfile, err := parser.WriteCSVFile(data, &badOutfilePath)
	assert.Error(t, err)
	assert.Nil(t, outfile)

	// Expect success for a good path.
	goodOutfilePath := "../testdata/testoutput.csv"
	outfile, err = parser.WriteCSVFile(data, &goodOutfilePath)
	assert.NoError(t, err)
	assert.NotNil(t, outfile)

	// Remove the file we just created.
	os.Remove(outfile.Name())
}

func TestTruncateColumnHeaders(t *testing.T) {
	testcases := []struct {
		description string
		input       []string
		expected    []string
	}{
		{
			description: "empty slice",
			input:       []string{},
			expected:    nil,
		},
		{
			description: "slice with one element",
			input:       []string{"single_header"},
			expected:    []string{"single_header"},
		},
		{
			description: "simple",
			input:       []string{"data_test_one", "data_test_two", "data_three"},
			expected:    []string{"test_one", "test_two", "three"},
		},
		{
			description: "long",
			input: []string{
				"data_one_two_three_four_after_header1",
				"data_one_two_three_four_after_header2",
				"data_one_two_three_four_after_header3",
				"data_one_two_three_four_after_header4",
				"data_one_two_three_four_after_header5",
				"data_one_two_three_four_before_header1",
				"data_one_two_three_four_before_header2",
				"data_one_two_three_four_before_header3",
				"data_one_two_three_four_before_header4",
				"data_one_two_three_four_before_header5",
				"data_one_two_three_four_header6",
				"data_one_two_three_four_events_header7",
				"data_one_two_three_four_events_header8",
				"data_one_two_three_four_header4",
			},
			expected: []string{
				"after_header1",
				"after_header2",
				"after_header3",
				"after_header4",
				"after_header5",
				"before_header1",
				"before_header2",
				"before_header3",
				"before_header4",
				"before_header5",
				"header6",
				"events_header7",
				"events_header8",
				"header4",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := parser.TruncateColumnHeaders(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
