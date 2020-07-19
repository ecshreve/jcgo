package parser

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/samsarahq/go/oops"

	oo "github.com/ecshreve/jcgo/internal/object"
)

// Parser is a representation of a JSON to CSV parsing session.
type Parser struct {
	Raw             map[string]interface{}
	RootObj         oo.Object
	ParsedData      [][]string
	TruncateHeaders bool
	InfilePath      *string
	OutfilePath     *string
	Outfile         *os.File
}

// NewParser returns a new instance of a Parser.
func NewParser(truncateHeaders bool) *Parser {
	return &Parser{
		TruncateHeaders: truncateHeaders,
	}
}

// BuildRootObj sets the Parser's RootObj to field to the Object representation
// of the map defined in the Parser's Raw field. It returns an error if unable
// to build the Object.
func (p *Parser) BuildRootObj() error {
	obj, err := oo.FromInterface("", p.Raw)
	if err != nil {
		return oops.Wrapf(err, "unable to build Object from interface")
	}

	p.RootObj = obj
	return nil
}

// Parse builds a 2d slice of strings from the Parser's RootObj, returns an
// error if unsuccessful.
func (p *Parser) Parse() error {
	if p.RootObj == nil {
		return oops.Errorf("no root object defined on Parser")
	}

	// Parse the Object into a [][]string.
	parsed, err := p.RootObj.Parse()
	if err != nil {
		return oops.Wrapf(err, "unable to parse Object")
	}

	p.ParsedData = parsed
	return nil
}

// ReadJSONFile reads the JSON file at the given path into the Parser's Raw map.
// It returns an error if reading the JSON file was unsuccessful.
func (p *Parser) ReadJSONFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return oops.Wrapf(err, "unable to open file %s", path)
	}
	defer file.Close()

	// Read the file into a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return oops.Wrapf(err, "unable to read file %s to byte array", file.Name())
	}

	// Unmarshall the byte array into a map.
	var result map[string]interface{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return oops.Wrapf(err, "unable to unmarshal byte array to map")
	}

	// Store the map in the Parser.
	p.Raw = result
	return nil
}

// WriteCSVFile writes the given 2d slice of strings to a CSV file, returns an
// error if unsuccessful.
//
// This function treats the first row in the input as headers for the CSV file.
func (p *Parser) WriteCSVFile(data [][]string) error {
	// Check if an OutfilePath is already defined, if not set it to the default.
	if p.OutfilePath == nil {
		p.OutfilePath = GetDefaultOutfilePath()
	}

	// Create the output file.
	file, err := os.Create(*p.OutfilePath)
	if err != nil {
		return oops.Wrapf(err, "unable to create output file for path: %v \n %v", *p.OutfilePath, file)
	}
	defer file.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// If the Parser is configured to do so, remove the longest common prefix
	// among all of the header strings.
	if p.TruncateHeaders {
		data[0] = TruncateColumnHeaders(data[0])
	}

	// Write each row in the data to a CSV file at the given path.
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return oops.Wrapf(err, "unable to write value: %+v to file: %s", value, file.Name())
		}
	}

	// Store a reference to the output file in the Parser.
	p.Outfile = file
	return nil
}
