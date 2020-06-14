package parser

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	oo "github.com/ecshreve/jcgo/internal/object"
	"github.com/samsarahq/go/oops"
)

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
func Convert(path string) (*os.File, error) {
	// Read the JSON file into a map.
	raw, err := ReadJSONFile(path)
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

	// Build the path for the output file.
	splitPath := strings.Split(path, ".")
	splitPath[len(splitPath)-1] = "output"
	splitPath = append(splitPath, "csv")
	outputPath := strings.Join(splitPath, ".")

	// Write the parsed data to CSV file.
	file, err := WriteCSVFile(parsed, outputPath)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to write to csv file")
	}

	return file, nil
}
