package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/samsarahq/go/oops"
)

// ReadJSONFile returns a pointer to the map representation of the JSON file at
// the given path or an error if reading the JSON file was unsuccessful.
func ReadJSONFile(path string) (*map[string]interface{}, error) {
	// Check that the given file is a JSON file.
	ext := filepath.Ext(path)
	if ext != ".json" {
		return nil, oops.Errorf("input file must be a JSON file: %s", path)
	}

	// Open the file specified by path.
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
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return nil, oops.Wrapf(err, "unable to unmarshal byte array to map")
	}

	return &result, nil
}

// WriteCSVFile writes the given 2d slice of strings to a CSV file at the given
// path. It returns a pointer to the output file, or an error if unsuccessful.
//
// If no path is provided, then a default filename is generated. This function
// treats the first row in the data argument as the headers for  the CSV file.
func WriteCSVFile(data [][]string, path *string) (*os.File, error) {
	var outfilePath *string
	// If no path is provided create a default output filename.
	if path == nil {
		outfilePath = GetDefaultOutfilePath()
	} else {
		// Check that the given output file is a CSV file.
		ext := filepath.Ext(*path)
		if ext != ".csv" {
			return nil, oops.Errorf("output file must be a CSV file: %s", *path)
		}
		outfilePath = path
	}

	// Create the output file.
	file, err := os.Create(*outfilePath)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create output file for path: %s", *outfilePath)
	}
	defer file.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write each row in the data to a CSV file at the given path.
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to write value: %+v to file: %s", value, file.Name())
		}
	}

	return file, nil
}

// TruncateColumnHeaders returns a slice of strings with the longest common
// prefix among all the elements removed from each.
func TruncateColumnHeaders(headers []string) []string {
	if len(headers) == 0 {
		return nil
	}

	if len(headers) == 1 {
		return headers
	}

	// Split each header into a slice of prefix tokens.
	var split [][]string
	for _, header := range headers {
		split = append(split, strings.Split(header, "_"))
	}

	// Sort by length so we can easily get the shortest header slice.
	sort.SliceStable(split, func(i, j int) bool {
		return len(split[i]) < len(split[j])
	})
	shortestHeader := split[0]

	// Keep a slice of the common prefix tokens.
	var validPrefs []string

	// Iterate through each header slice and break out when we encounter a
	// prefix token that doesn't match corresponding prefix in the shortest
	// headers slice.
	valid := true
	for i := 0; i < len(shortestHeader) && valid; i++ {
		for _, row := range split {
			if row[i] != split[0][i] {
				valid = false
				break
			}
		}
		if valid {
			validPrefs = append(validPrefs, split[0][i])
		}
	}

	// Join the slice of valid prefixes back into a string and add a trailing
	// underscore so we can remove the common prefix from each header.
	longestPrefix := strings.Join(validPrefs, "_") + "_"

	// Remove the longest common prefix from each header.
	for i, header := range headers {
		headers[i] = strings.Replace(header, longestPrefix, "", 1)
	}

	return headers
}

// GetDefaultOutfilePath returns the default file path for an output file.
//
// The default directory for the output file is the root directory of the
// module. The filename is of the form `data_<seconds_epoch>.output.csv`.
func GetDefaultOutfilePath() *string {
	timeNowMs := time.Now().Unix()
	outFilePath := fmt.Sprintf("data_%d.output.csv", timeNowMs)
	return &outFilePath
}
