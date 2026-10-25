package cli

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"csvstat/internal/csvdata"
)

// Exit codes. Only cmd/csvstat calls os.Exit with the value Run returns.
const (
	exitOK          = 0
	exitOperational = 1
	exitUsage       = 2
	exitRejected    = 3
)

// usageLine shows both command forms on one line of standard error.
const usageLine = "usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>"

// Run executes csvstat with args, which excludes the program name, and returns
// the process exit code. Stdout is written only on success; diagnostics go to
// stderr. Only cmd/csvstat calls os.Exit with this value.
func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	opts, err := ParseArgs(args)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v; %s\n", err, usageLine)
		return exitUsage
	}

	f, err := os.Open(opts.File)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitOperational
	}
	defer f.Close()

	table, err := csvdata.Read(f)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %s: %v\n", opts.File, err)
		return readFailureExit(err)
	}
	csvdata.Trim(table)
	if err := csvdata.Validate(table); err != nil {
		fmt.Fprintf(stderr, "csvstat: %s: %v\n", opts.File, err)
		return exitRejected
	}

	var rows [][]string
	switch {
	case opts.FreqSet:
		rows, err = freqRows(table, opts)
		if err != nil {
			fmt.Fprintf(stderr, "csvstat: %s: %v\n", opts.File, err)
			return exitRejected
		}
	case opts.ColumnSet:
		rows, err = columnRows(table, opts.Column)
		if err != nil {
			fmt.Fprintf(stderr, "csvstat: %s: %v\n", opts.File, err)
			return exitRejected
		}
	default:
		rows = profileRows(table)
	}
	return writeRows(stdout, stderr, rows)
}

// writeRows encodes rows as RFC 4180 CSV with \n line endings. The whole output is built in memory
// first, so stdout receives nothing unless encoding succeeds.
func writeRows(stdout io.Writer, stderr io.Writer, rows [][]string) int {
	var buf bytes.Buffer
	if err := csv.NewWriter(&buf).WriteAll(rows); err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitOperational
	}
	if _, err := stdout.Write(buf.Bytes()); err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitOperational
	}
	return exitOK
}

// readFailureExit maps a Read error to an exit code: a CSV syntax error is
// rejected input, while any other failure to read the file is operational.
func readFailureExit(err error) int {
	var parseErr *csv.ParseError
	if errors.As(err, &parseErr) {
		return exitRejected
	}
	return exitOperational
}
