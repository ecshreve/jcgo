package main

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
)

func TestJCGOEndToEnd(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	testcases := []struct {
		description string
		infilePath  string
		outfilePath string
	}{
		{
			description: "valid infilePath, valid outfilePath, no extraArg expect success",
			infilePath:  "testdata/json1.json",
			outfilePath: "testdata/json1.output.csv",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			// Save the existing command line args, and reset them at the end of
			// this testcase.
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// Temporarily set command line args before running main func.
			os.Args = []string{"dummy_process_name"}
			if len(testcase.infilePath) > 0 {
				os.Args = append(os.Args, testcase.infilePath)
			}
			if len(testcase.outfilePath) > 0 {
				os.Args = append(os.Args, testcase.outfilePath)
			}

			// Call the main function with the testcases's command line args.
			main()

			// Open the expected output file and expect no error.
			file, err := os.Open(testcase.outfilePath)
			assert.NoError(t, err)

			// Read data from the output file and expect no error.
			fileReader := csv.NewReader(file)
			data, err := fileReader.ReadAll()
			assert.NoError(t, err)

			// Snapshot the data read from the output file.
			snap.Snapshot(testcase.description, data)
		})
	}
}
