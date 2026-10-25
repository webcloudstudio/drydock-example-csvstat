// Package cli parses the csvstat command line and dispatches it. Run takes its
// writers as arguments so tests drive it without touching the process streams.
package cli

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"

	"csvstat/internal/csvdata"
	"csvstat/internal/profile"
)

// Exit codes returned by Run. Only cmd/csvstat passes them to os.Exit.
const (
	ExitOK          = 0 // success
	ExitOperational = 1 // the file cannot be opened or read
	ExitUsage       = 2 // invalid command line; checked before the file is opened
	ExitRejected    = 3 // rejected input: empty file, bad header, ragged record, unknown column
)

// defaultTop is the number of frequent values listed when --freq is given without --top.
const defaultTop = 5

// options holds the validated command line.
type options struct {
	file      string
	freq      string
	column    string
	top       int
	columnSet bool // --column was given, even when its value is empty
	freqSet   bool // --freq was given, even when its value is empty
}

// Run executes the csvstat command with args (excluding the program name) and
// returns the exit code. A failing run writes nothing to stdout and exactly one
// diagnostic line to stderr.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, err := parse(args, stdout)
	if errors.Is(err, flag.ErrHelp) {
		return ExitOK
	}
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return ExitUsage
	}
	return report(opts, stdout, stderr)
}

// parse reads and validates the command line. Help requests are written to help.
func parse(args []string, help io.Writer) (options, error) {
	var opts options
	fs := newFlagSet(&opts)
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fs.SetOutput(help)
			fs.Usage()
		}
		return options{}, err
	}

	set := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { set[f.Name] = true })

	if fs.NArg() != 1 {
		return options{}, fmt.Errorf("expected exactly one FILE argument, got %d", fs.NArg())
	}
	opts.file = fs.Arg(0)
	opts.columnSet = set["column"]
	opts.freqSet = set["freq"]
	if err := validate(opts, set); err != nil {
		return options{}, err
	}
	return opts, nil
}

// newFlagSet declares the options. Its output is discarded so that parse errors
// produce only the single diagnostic line written by Run.
func newFlagSet(opts *options) *flag.FlagSet {
	fs := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "usage: csvstat [--column <name> | --freq <name> [--top <n>]] <file>")
		fs.PrintDefaults()
	}
	fs.StringVar(&opts.freq, "freq", "", "list the most frequent values of the named column")
	fs.StringVar(&opts.column, "column", "", "restrict output to the named column")
	fs.IntVar(&opts.top, "top", defaultTop, "number of frequent values to list (positive; requires --freq)")
	return fs
}

// validate applies the cross-option rules. set reports which flags were given
// explicitly, so an unset --top is distinguishable from --top 0.
func validate(opts options, set map[string]bool) error {
	if set["top"] && !set["freq"] {
		return errors.New("--top requires --freq")
	}
	if set["top"] && opts.top <= 0 {
		return fmt.Errorf("--top must be a positive integer, got %d", opts.top)
	}
	if set["freq"] && set["column"] {
		return errors.New("--column cannot be used with --freq")
	}
	return nil
}

// report reads and validates the file, then writes the profile. The named column
// of --column or --freq must be in the header, checked after validation. Frequency
// listing is built by a later story, so --freq exits 0 with nothing on stdout once
// its column is found.
func report(opts options, stdout, stderr io.Writer) int {
	table, err := csvdata.Read(opts.file)
	if err != nil {
		return fail(stderr, opts.file, err)
	}
	if err := csvdata.Validate(table); err != nil {
		return fail(stderr, opts.file, err)
	}

	rows := profile.Rows(table)
	if opts.columnSet || opts.freqSet {
		name := opts.column
		if opts.freqSet {
			name = opts.freq
		}
		index, found := csvdata.Column(table, name)
		if !found {
			return fail(stderr, opts.file, &csvdata.RejectError{
				Reason: fmt.Sprintf("column %q is not in the header", name),
			})
		}
		if opts.freqSet {
			return ExitOK
		}
		rows = rows[index : index+1]
	}

	if err := writeCSV(stdout, rows); err != nil {
		return fail(stderr, opts.file, err)
	}
	return ExitOK
}

// writeCSV writes the profile header and rows as CSV with \n line endings.
func writeCSV(stdout io.Writer, rows [][]string) error {
	w := csv.NewWriter(stdout)
	return w.WriteAll(append([][]string{profile.Header}, rows...))
}

// fail writes the single diagnostic line for a file that cannot be profiled and
// returns its exit code. Content that was read but rejected is ExitRejected; any
// other failure to read the file is ExitOperational.
func fail(stderr io.Writer, file string, err error) int {
	fmt.Fprintf(stderr, "csvstat: %s: %v\n", file, err)
	var rejected *csvdata.RejectError
	if errors.As(err, &rejected) {
		return ExitRejected
	}
	return ExitOperational
}
