package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ecshreve/jcgo/parser"
)

func main() {
	args := os.Args

	// Require an argument.
	if len(args) <= 1 {
		log.Fatalf("please provide a file")
	}

	// Require that argument is a path to a .json file.
	path := args[1]
	ext := filepath.Ext(path)
	if ext != ".json" {
		log.Fatalf("please provide a .json file")
	}

	// Open the file.
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successfully opened file %s\n", path)
	defer file.Close()

	// Read the file into a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshall the byte array into a map.
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	transformed := parser.Transform(result)
	parsed := parser.Parse(transformed)

	outfile, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successfully created file %s\n", outfile.Name())
	defer outfile.Close()

	writer := csv.NewWriter(outfile)
	defer writer.Flush()

	for _, value := range parsed {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
