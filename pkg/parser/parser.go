package parser

import (
	"os"

	"github.com/samsarahq/go/oops"

	oo "github.com/ecshreve/jcgo/internal/object"
)

// Parser is a representation of a JSON to CSV parsing session.
type Parser struct {
	Raw             *map[string]interface{}
	RootObj         oo.Object
	ParsedData      [][]string
	TruncateHeaders bool
	InfilePath      *string
	OutfilePath     *string
	Outfile         *os.File
}

// NewParser returns a new instance of a Parser.
func NewParser(truncateHeaders bool, infilePath string) *Parser {
	return &Parser{
		TruncateHeaders: truncateHeaders,
		InfilePath:      &infilePath,
	}
}

// ConvertJSONFile converts a JSON file at the given path to a CSV file, and
// returns a pointer to the newly created file, or an error if unsuccessful.
func ConvertJSONFile(path string) (*os.File, error) {
	pp := NewParser(true, path)

	err := pp.readJSONFile()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to read json file: %s", path)
	}

	err = pp.BuildRootObj()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build root object for map: %v", pp.Raw)
	}

	err = pp.Parse()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to parse root object: %v", pp.RootObj)
	}

	err = pp.writeCSVFile()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to write data to csv file, data: %v", pp.ParsedData)
	}

	return pp.Outfile, nil
}

// BuildRootObj sets the Parser's RootObj field to the Object representation of
// the map defined in the Parser's Raw field. It returns an error if unable to
// build the Object.
func (p *Parser) BuildRootObj() error {
	obj, err := oo.FromInterface("", *p.Raw)
	if err != nil {
		return oops.Wrapf(err, "unable to build Object from interface")
	}

	p.RootObj = obj
	return nil
}

// Parse sets the Parser's ParsedData field to a 2d slice of strings built from
// the Parser's RootObj, returns an error if unsuccessful.
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

// readJSONFile reads the JSON file specified by the Parser's InfilePath field
// and stores the resulting map in the Parser's Raw field. Returns an error if
// reading the file was unnsuccessful.
func (p *Parser) readJSONFile() error {
	raw, err := ReadJSONFile(*p.InfilePath)
	if err != nil {
		return oops.Wrapf(err, "unable to read json file: %s", *p.InfilePath)
	}

	// Store a pointer to the map in the Parser.
	p.Raw = raw

	return nil
}

// writeCSVFile writes the data in the Parser's ParsedData field to the CSV file
// defined in the Parser's OutfilePath field. Returns an error if unsuccessful.
//
// If the Parser's TruncateHeaders field is set to true then headers in the
// first row of the Parser's ParsedData field are truncated prior to writing the
// CSV file.
func (p *Parser) writeCSVFile() error {
	// Check if an OutfilePath is already defined, if not set it to the default.
	if p.OutfilePath == nil {
		p.OutfilePath = GetDefaultOutfilePath()
	}

	// If the Parser is configured to do so, remove the longest common prefix
	// among all of the header strings.
	if p.TruncateHeaders {
		p.ParsedData[0] = TruncateColumnHeaders(p.ParsedData[0])
	}

	// Write the Parser's ParsedData to a csv file.
	outfile, err := WriteCSVFile(p.ParsedData, p.OutfilePath)
	if err != nil {
		return oops.Wrapf(err, "unable to write csv file: %s", *p.OutfilePath)
	}

	// Store a reference to the output file in the Parser.
	p.Outfile = outfile

	return nil
}
