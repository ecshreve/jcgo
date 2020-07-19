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

	pp := parser.NewParser(true)

	err := pp.ReadJSONFile(os.Args[1])
	if err != nil {
		log.Fatalf("error reading json file: %v", err)
	}

	err = pp.BuildRootObj()
	if err != nil {
		log.Fatalf("error building root object: %v", err)
	}

	err = pp.Parse()
	if err != nil {
		log.Fatalf("error parsing root object: %v", err)
	}

	err = pp.WriteCSVFile(pp.ParsedData)
	if err != nil {
		log.Fatalf("error writing json file: %v", err)
	}

	log.Printf("generated csv file: %v\n", pp.Outfile.Name())
}
