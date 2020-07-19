package main

import (
	"log"
	"os"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please provide an input file")
	}

	if len(os.Args) > 3 {
		log.Fatal("too many command line arguments")
	}

	infilePath := &os.Args[1]

	var outfilePath *string
	if len(os.Args) == 3 {
		outfilePath = &os.Args[2]
	}

	outfile, err := parser.ConvertJSONFile(infilePath, outfilePath)
	if err != nil {
		log.Fatalf("error converting json file: %v", err)
	}

	log.Printf("generated csv file: %v\n", outfile.Name())
}
