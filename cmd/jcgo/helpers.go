package main

import (
	"bytes"
	"flag"
	"log"

	"github.com/samsarahq/go/oops"

	"github.com/ecshreve/jcgo/internal/parser"
)

// parseArgs returns a parser.Config for the given slice of command line args,
// along with any output message or error that result from parsing.
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

// handleParseError returns an exit status code and an error for the given
// output string and error that result from parsing command line args.
func handleParseError(output string, err error) (int, error) {
	// Parsing should result in printing the help message to the console, and
	// exiting the program. But we don't want to differentiate this type of exit
	// from a generic Fatal exit, so we give it a status code of 2.
	if err == flag.ErrHelp {
		log.Println(output)
		return 2, nil
	}

	// If parsing resulted in an error that wasn't a help error, then return it
	// along with an exit status code of 1 so we can Fatal exit.
	if err != nil {
		return 1, oops.Wrapf(err, "unable to parse flags -- %s", output)
	}

	// Parsing was successful, so return the exit status code of 0 indicating
	// success, and no error.
	return 0, nil
}
