package parser

import (
	"bytes"
	"flag"
	"log"
	"strings"

	"github.com/samsarahq/go/oops"
)

// ParseArgs returns a Config for the given slice of command line args,
// along with any output message or error that result from parsing.
func ParseArgs(args []string) (*Config, string, error) {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var cfg Config
	flags.StringVar(&cfg.Infile, "infile", "", "specify input file")
	flags.StringVar(&cfg.Outfile, "outfile", "", "specify output file")

	err := flags.Parse(args[1:])
	if err != nil {
		return nil, buf.String(), err
	}

	cfg.Args = flags.Args()
	return &cfg, buf.String(), nil
}

// HandleParseError returns an exit status code and an error for the given
// output string and error that result from parsing command line args.
func HandleParseError(output string, err error) (int, error) {
	// Parsing should result in printing the help message to the console, and
	// exiting the program. But we want to differentiate this type of exit
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

// SetDefaultOutfile sets the Outfile field on the given Config to be the
// outfile path. The default path for an outfile is defined as
// [INFILE_BASE].output.csv. Returns an error if an Outfile already exists on
// the given Config.
func (cfg *Config) SetDefaultOutfile() error {
	if len(cfg.Outfile) > 0 {
		return oops.Errorf("outfile already exists on Config: %v", cfg)
	}

	// Build the path for the output file.
	splitPath := strings.Split(cfg.Infile, ".")
	splitPath[len(splitPath)-1] = "output"
	splitPath = append(splitPath, "csv")
	outputPath := strings.Join(splitPath, ".")
	cfg.Outfile = outputPath

	return nil
}
