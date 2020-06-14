package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/ecshreve/jcgo/internal/parser"
	"github.com/samsarahq/go/oops"
)

func parseArgs(args []string) (*parser.Config, string, error) {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var cfg parser.Config
	flags.StringVar(&cfg.Infile, "infile", "", "specify input file")

	err := flags.Parse(args[1:])
	if err != nil {
		return nil, buf.String(), err
	}

	cfg.Args = flags.Args()
	return &cfg, buf.String(), nil
}

func handleParseError(output string, err error) (int, error) {
	if err == flag.ErrHelp {
		log.Println(output)
		return 2, nil
	}

	if err != nil {
		return 1, oops.Wrapf(err, "unable to parse flags -- %s", output)
	}

	return 0, nil
}

func main() {
	cfg, output, err := parseArgs(os.Args)

	// Handle any errors that resulted from parsing flags.
	exitCode, err := handleParseError(output, err)
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
