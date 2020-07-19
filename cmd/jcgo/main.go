package main

import (
	"log"
	"os"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide an input file")
	}

	outfile, err := parser.ConvertJSONFile(os.Args[1])
	if err != nil {
		log.Fatalf("error converting json file: %v", err)
	}

	log.Printf("generated csv file: %v\n", outfile.Name())
}
