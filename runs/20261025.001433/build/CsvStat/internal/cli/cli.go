// Package cli parses csvstat arguments, prints usage, and maps outcomes to exit codes.
package cli

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"csvstat/internal/csvio"
)

const (
	exitOK     = 0
	exitRead   = 1
	exitUsage  = 2
	exitReject = 3
)

const usageText = `usage:
  csvstat [--column <name>] <file>
  csvstat --freq <name> [--top <n>] <file>
  csvstat --help

options:
  --column <name>  profile only the named column
  --freq <name>    list the most frequent values of the named column
  --top <n>        with --freq, list at most n values (default 5)
  --help           show this usage text
`

// Run executes csvstat with args (the program name excluded) and returns the process exit code.
// Diagnostics go to stderr; stdout receives only usage text or profile output.
func Run(args []string, stdout, stderr io.Writer) int {
	return run(args, stdout, stderr, profileMode{}, frequencyMode{})
}

// run dispatches a parsed request to the profile or frequency mode. Output is buffered and written
// only after the whole file has been read and validated.
func run(args []string, stdout, stderr io.Writer, profiler Profiler, frequencies FrequencyLister) int {
	req, err := parse(args)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		usage(stderr)
		return exitUsage
	}
	if req.Help {
		usage(stdout)
		return exitOK
	}

	table, err := csvio.Load(req.Path)
	if err != nil {
		return fail(stderr, err)
	}

	var rows [][]string
	if req.Mode == ModeFrequency {
		column, ok := table.Index(req.Column)
		if !ok {
			return fail(stderr, absentColumn(req))
		}
		rows = frequencies.Frequencies(table, column, req.Top)
	} else {
		columns, err := selectColumns(table, req)
		if err != nil {
			return fail(stderr, err)
		}
		rows = profiler.Profile(table, columns)
	}

	var buf bytes.Buffer
	if err := csv.NewWriter(&buf).WriteAll(rows); err != nil {
		fmt.Fprintf(stderr, "csvstat: write output: %v\n", err)
		return exitRead
	}
	if _, err := stdout.Write(buf.Bytes()); err != nil {
		fmt.Fprintf(stderr, "csvstat: write output: %v\n", err)
		return exitRead
	}
	return exitOK
}

// selectColumns returns the column positions the profile mode reports.
func selectColumns(t *csvio.Table, req Request) ([]int, error) {
	if req.AllColumns {
		columns := make([]int, len(t.Header))
		for i := range columns {
			columns[i] = i
		}
		return columns, nil
	}
	column, ok := t.Index(req.Column)
	if !ok {
		return nil, absentColumn(req)
	}
	return []int{column}, nil
}

// absentColumn is the rejection for a --column or --freq name that is not in the header.
func absentColumn(req Request) error {
	return &csvio.RejectError{Path: req.Path, Reason: fmt.Sprintf("no column named %q", req.Column)}
}

// fail writes the diagnostic for err to stderr and returns its exit code.
func fail(stderr io.Writer, err error) int {
	fmt.Fprintf(stderr, "csvstat: %v\n", err)
	var reject *csvio.RejectError
	if errors.As(err, &reject) {
		return exitReject
	}
	return exitRead
}

func usage(w io.Writer) {
	io.WriteString(w, usageText)
}
