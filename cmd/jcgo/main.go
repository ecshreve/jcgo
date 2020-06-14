package main

import (
	"log"
	"os"

	"github.com/ecshreve/jcgo/internal/parser"
)

func main() {
	cfg, output, err := parser.ParseArgs(os.Args)
	// Handle any errors that resulted from parsing flags.
	exitCode, err := parser.HandleParseError(output, err)
	if err != nil {
		log.Print(err)
		os.Exit(exitCode)
	}

	// Validate the parser.Config that results from parsing flags.
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the JSON file to a CSV file.
	file, err := parser.Convert(cfg.Infile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successfully created %s", file.Name())
}
