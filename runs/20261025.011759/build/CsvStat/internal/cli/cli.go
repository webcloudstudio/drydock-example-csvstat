// Package cli parses csvstat arguments, maps outcomes to exit codes, and drives the profiler.
package cli

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"

	"csvstat/internal/csvio"
	"csvstat/internal/output"
	"csvstat/internal/profile"
)

// Exit codes returned by Run.
const (
	exitOK       = 0
	exitFailure  = 1
	exitUsage    = 2
	exitRejected = 3
)

const (
	defaultTop = 5
	usageLine  = "usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>"
)

// mode selects which output the command writes.
type mode int

const (
	modeProfile mode = iota
	modeColumn
	modeFrequency
)

// request is a validated command line.
type request struct {
	mode   mode
	file   string
	column string
	top    int
}

// Run executes csvstat with args (excluding the program name). Results go to stdout,
// diagnostics go to stderr, and the returned value is the process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	req, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(stderr, usageLine)
		return exitUsage
	}

	table, err := csvio.Read(req.file)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		if errors.Is(err, csvio.ErrRejected) {
			return exitRejected
		}
		return exitFailure
	}

	// Everything is validated and computed before the first byte reaches stdout.
	var buf bytes.Buffer
	switch req.mode {
	case modeFrequency:
		idx, ok := columnIndex(table, req.column)
		if !ok {
			return columnMissing(stderr, req)
		}
		freqs := profile.Frequencies(columnValues(table, idx))
		if len(freqs) > req.top {
			freqs = freqs[:req.top]
		}
		err = output.WriteFrequencies(&buf, freqs)
	case modeColumn:
		idx, ok := columnIndex(table, req.column)
		if !ok {
			return columnMissing(stderr, req)
		}
		col := profile.ProfileColumn(table.Header[idx], columnValues(table, idx))
		err = output.WriteProfile(&buf, []profile.Column{col})
	default:
		cols := make([]profile.Column, len(table.Header))
		for i, name := range table.Header {
			cols[i] = profile.ProfileColumn(name, columnValues(table, i))
		}
		err = output.WriteProfile(&buf, cols)
	}
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: write output: %v\n", err)
		return exitFailure
	}
	if _, err := stdout.Write(buf.Bytes()); err != nil {
		fmt.Fprintf(stderr, "csvstat: write output: %v\n", err)
		return exitFailure
	}
	return exitOK
}

// parseArgs validates the command line. Options precede the single file argument.
func parseArgs(args []string) (request, error) {
	fs := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	column := fs.String("column", "", "")
	freq := fs.String("freq", "", "")
	top := fs.String("top", "", "")
	if err := fs.Parse(args); err != nil {
		return request{}, err
	}
	if fs.NArg() != 1 {
		return request{}, errors.New("expected exactly one file")
	}

	set := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { set[f.Name] = true })
	if set["column"] && set["freq"] {
		return request{}, errors.New("--column and --freq cannot be combined")
	}
	if set["top"] && !set["freq"] {
		return request{}, errors.New("--top requires --freq")
	}

	req := request{mode: modeProfile, file: fs.Arg(0), top: defaultTop}
	switch {
	case set["freq"]:
		req.mode = modeFrequency
		req.column = *freq
	case set["column"]:
		req.mode = modeColumn
		req.column = *column
	}
	if set["top"] {
		n, err := strconv.Atoi(*top)
		if err != nil || n < 1 {
			return request{}, fmt.Errorf("--top must be a positive integer, got %q", *top)
		}
		req.top = n
	}
	return req, nil
}

// columnMissing reports a column name absent from the header, which is a rejected request.
func columnMissing(stderr io.Writer, req request) int {
	fmt.Fprintf(stderr, "csvstat: no column named %q in %s\n", req.column, req.file)
	return exitRejected
}

// columnIndex finds the header position of a name by exact match against trimmed names.
func columnIndex(table *csvio.Table, name string) (int, bool) {
	for i, h := range table.Header {
		if h == name {
			return i, true
		}
	}
	return 0, false
}

// columnValues returns the values of column idx in record order.
func columnValues(table *csvio.Table, idx int) []string {
	values := make([]string, len(table.Rows))
	for r, row := range table.Rows {
		values[r] = row[idx]
	}
	return values
}
