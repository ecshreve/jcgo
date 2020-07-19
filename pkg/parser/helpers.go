package parser

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

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
