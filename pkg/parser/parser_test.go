package parser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func TestReadJSONFile(t *testing.T) {
	pp := parser.NewParser(true)

	// Expect error if file doesn't exist.
	err := pp.ReadJSONFile("../testdata/nonexistentfile.json")
	assert.Error(t, err)
	assert.Nil(t, pp.Raw)

	// Expect error if file is malformed.
	err = pp.ReadJSONFile("../testdata/jsontest_bad.json")
	assert.Error(t, err)
	assert.Nil(t, pp.Raw)

	// Verify we can read a valid JSON file.
	err = pp.ReadJSONFile("../testdata/jsontest.json")
	assert.NoError(t, err)
	assert.NotNil(t, pp.Raw)
}

func TestWriteJSONFile(t *testing.T) {
	pp := parser.NewParser(true)

	// Verify we can write valid data to CSV.
	data := [][]string{
		{"one", "two"},
		{"one_one", "two_two"},
	}

	// Expect error for a bad path.
	badOutfilePath := "../testdata/nonexistentdirectory/testoutput.xls"
	pp.OutfilePath = &badOutfilePath
	err := pp.WriteCSVFile(data)
	assert.Error(t, err)
	assert.Nil(t, pp.Outfile)

	// Expect success for a good path.
	goodOutfilePath := "../testdata/testoutput.csv"
	pp.OutfilePath = &goodOutfilePath
	err = pp.WriteCSVFile(data)
	assert.NoError(t, err)
	assert.NotNil(t, pp.Outfile)

	// Remove the file we just created.
	os.Remove(goodOutfilePath)
}
