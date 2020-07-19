package main

// func TestEndToEndSuccess(t *testing.T) {
// 	snap := snapshotter.New(t)
// 	defer snap.Verify()

// 	testcases := []struct {
// 		description string
// 		infile      string
// 		outfile     string
// 	}{
// 		{
// 			description: "basic",
// 			infile:      "testdata/json1.json",
// 			outfile:     "testdata/json1.output.csv",
// 		},
// 	}

// 	for _, testcase := range testcases {
// 		t.Run(testcase.description, func(t *testing.T) {
// 			// Save the existing command line args, and reset them at the end of
// 			// this testcase.
// 			oldArgs := os.Args
// 			defer func() { os.Args = oldArgs }()

// 			// Temporarily set command line args before running main func.
// 			os.Args = []string{"dummy_prog_name", fmt.Sprintf("--infile=%s", testcase.infile)}
// 			main()

// 			// Open the expected output file and expect no error.
// 			file, err := os.Open(testcase.outfile)
// 			assert.NoError(t, err)

// 			// Read data from the output file and expect no error.
// 			fileReader := csv.NewReader(file)
// 			data, err := fileReader.ReadAll()
// 			assert.NoError(t, err)

// 			// Snapshot the data read from the output file.
// 			snap.Snapshot(testcase.description, data)
// 		})
// 	}
// }
