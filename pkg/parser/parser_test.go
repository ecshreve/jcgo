package parser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func TestConvertJSONFile(t *testing.T) {
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
			outfile, err := parser.ConvertJSONFile(testcase.infilePath)
			if testcase.expectError {
				assert.Error(t, err)
				assert.Nil(t, outfile)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, outfile)
				os.Remove(outfile.Name())
			}
		})
	}
}
