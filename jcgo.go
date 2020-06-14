package main

import (
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

	file, err := parser.Convert(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successfully created %s", file.Name())
}
