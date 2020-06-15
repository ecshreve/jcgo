package parser

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/samsarahq/go/oops"

	"github.com/ecshreve/jcgo/internal/helpers"
	oo "github.com/ecshreve/jcgo/internal/object"
)

// Config represents various configuration parameters for one JSON to CSV
// parsing session.
type Config struct {
	Infile  string
	Outfile string

	Args []string
}

// Validate returns an error if the Config is invalid. A Config can be invalid
// for the following reasons:
//  - No input file provided.
//  - Non-JSON input file provided.
//  - Non-CSV output file provided.
//
// TODO
//  - is it possible for the Config to not have an outfile here? how should that
//    be handled?
//  - what happens if the Config's outfile already exists?
func (cfg *Config) Validate() error {
	if len(cfg.Infile) == 0 {
		return oops.Errorf("please provide a file with the --infile flag")
	}

	ext := filepath.Ext(cfg.Infile)
	if ext != ".json" {
		return oops.Errorf("infile mush be a .json file -- infile: %s", cfg.Infile)
	}

	ext = filepath.Ext(cfg.Outfile)
	if ext != ".csv" {
		return oops.Errorf("outfile mush be a .csv file -- outfile: %s", cfg.Outfile)
	}

	return nil
}

// ReadJSONFile reads the JSON file at the given path into a map, returns an
// error if unsuccessful.
func ReadJSONFile(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to open file %s", path)
	}
	defer file.Close()

	// Read the file into a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to read file %s to byte array", file.Name())
	}

	// Unmarshall the byte array into a map.
	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to unmarshal byte array to map")
	}

	return result, nil
}

// WriteCSVFile writes the given data to a CSV file at the given path, returns
// an error if unsuccessful.
func WriteCSVFile(data [][]string, path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create ouput file")
	}
	defer file.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Truncate header strings.
	data[0] = helpers.TruncateColumnHeaders(data[0])

	// Write each row in our data to the CSV file.
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to value: %+v to file: %s", value, file.Name())
		}
	}

	return file, nil
}

// Transform returns an Object for the given input map. It's meant to be called
// on the root map[string]interface{} that comes from Unmarshalling a JSON file
// into a map[string]interface{}.
func Transform(input map[string]interface{}) (oo.Object, error) {
	obj, err := oo.FromInterface("", input)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to get Object from interface")
	}

	return obj, nil
}

// Convert converts the given json file at the given path to a csv file, or
// returns an error if unable to convert the file.
func Convert(cfg *Config) (*os.File, error) {
	// Read the JSON file into a map.
	raw, err := ReadJSONFile(cfg.Infile)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to read json file")
	}

	// Transform the map into an Object.
	transformed, err := Transform(raw)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to convert to Object")
	}

	// Parse the Object into a [][]string.
	parsed, err := transformed.Parse()
	if err != nil {
		return nil, oops.Wrapf(err, "unable to parse Object")
	}

	// Write the parsed data to CSV file.
	file, err := WriteCSVFile(parsed, cfg.Outfile)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to write to csv file")
	}

	return file, nil
}
