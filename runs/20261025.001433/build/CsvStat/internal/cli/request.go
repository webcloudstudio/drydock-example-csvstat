package cli

import (
	"errors"
	"flag"
	"io"
	"math"
	"strconv"
)

// Mode selects which command form a request names.
type Mode int

const (
	// ModeProfile is the profile form: csvstat [--column <name>] <file>.
	ModeProfile Mode = iota
	// ModeFrequency is the frequency form: csvstat --freq <name> [--top <n>] <file>.
	ModeFrequency
)

// defaultTop is the --freq row limit used when --top is absent.
const defaultTop = 5

// Request is a well-formed csvstat command line.
type Request struct {
	Help       bool
	Mode       Mode
	Column     string // the name given to --column or --freq
	AllColumns bool   // profile mode without --column
	Top        int    // frequency mode row limit
	Path       string
}

// usageError marks a command line that matches neither command form.
type usageError struct {
	msg string
}

func (e *usageError) Error() string { return e.msg }

// parse turns args into a Request. Options precede exactly one file. Flag-package diagnostics are
// discarded here; the caller reports every usage error the same way.
func parse(args []string) (Request, error) {
	fs := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	column := fs.String("column", "", "")
	freq := fs.String("freq", "", "")
	top := fs.String("top", "", "")
	help := fs.Bool("help", false, "")
	if err := fs.Parse(args); err != nil {
		return Request{}, &usageError{msg: err.Error()}
	}
	if *help {
		return Request{Help: true}, nil
	}

	set := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { set[f.Name] = true })

	if fs.NArg() != 1 {
		return Request{}, &usageError{msg: "expected exactly one CSV file"}
	}
	if set["column"] && set["freq"] {
		return Request{}, &usageError{msg: "--column cannot be combined with --freq"}
	}
	if set["top"] && !set["freq"] {
		return Request{}, &usageError{msg: "--top requires --freq"}
	}

	req := Request{Path: fs.Arg(0), Top: defaultTop}
	if !set["freq"] {
		req.Mode = ModeProfile
		req.Column = *column
		req.AllColumns = !set["column"]
		return req, nil
	}

	req.Mode = ModeFrequency
	req.Column = *freq
	if set["top"] {
		n, err := parseTop(*top)
		if err != nil {
			return Request{}, &usageError{msg: err.Error()}
		}
		req.Top = n
	}
	return req, nil
}

// parseTop accepts a positive integer. A value too large for an int is unbounded, not rejected.
func parseTop(s string) (int, error) {
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, errors.New("--top must be a positive integer, got " + strconv.Quote(s))
		}
	}
	n, err := strconv.Atoi(s)
	if errors.Is(err, strconv.ErrRange) {
		return math.MaxInt, nil
	}
	if err != nil || n == 0 {
		return 0, errors.New("--top must be a positive integer, got " + strconv.Quote(s))
	}
	return n, nil
}
