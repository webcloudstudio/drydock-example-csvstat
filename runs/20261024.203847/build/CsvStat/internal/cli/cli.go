// Package cli implements the csvstat command line.
package cli

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"csvstat/internal/freq"
	"csvstat/internal/output"
	"csvstat/internal/profile"
	"csvstat/internal/reader"
)

const (
	exitOK       = 0
	exitFailure  = 1 // the file cannot be opened or read
	exitUsage    = 2 // the command line is malformed
	exitRejected = 3 // the file was read but rejected, or a named column is absent
	defaultTop   = 5
)

const (
	profileSynopsis = "csvstat [--column <name>] <file>"
	freqSynopsis    = "csvstat --freq <name> [--top <n>] <file>"
)

// request is a parsed command line.
type request struct {
	file      string
	column    string
	hasColumn bool
	freqName  string
	hasFreq   bool
	top       int
}

// Run executes csvstat with the given arguments and returns the process exit code.
// Standard output carries only profile or frequency CSV, and only on success.
// Diagnostics go to stderr.
func Run(args []string, stdout, stderr io.Writer) int {
	req, ok := parseArgs(args, stderr)
	if !ok {
		return exitUsage
	}

	f, err := os.Open(req.file)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitFailure
	}
	defer f.Close()

	table, err := reader.Read(f)
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %s: %v\n", req.file, err)
		var parseErr *csv.ParseError
		if errors.As(err, &parseErr) {
			return exitRejected
		}
		return exitFailure
	}

	table = reader.TrimTable(table)
	if err := reader.Validate(table); err != nil {
		fmt.Fprintf(stderr, "csvstat: %s: %v\n", req.file, err)
		return exitRejected
	}

	var out bytes.Buffer
	if req.hasFreq {
		j, found := columnIndex(table.Header, req.freqName)
		if !found {
			return absentColumn(stderr, req.file, req.freqName)
		}
		entries := freq.Top(columnValues(table, j), req.top)
		err = output.WriteFreq(&out, entries)
	} else {
		cols, found := profiledColumns(table, req)
		if !found {
			return absentColumn(stderr, req.file, req.column)
		}
		profiles := make([]profile.Column, 0, len(cols))
		for _, j := range cols {
			profiles = append(profiles, profile.Describe(table.Header[j], columnValues(table, j)))
		}
		err = output.WriteProfile(&out, profiles)
	}
	if err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitFailure
	}

	if _, err := stdout.Write(out.Bytes()); err != nil {
		fmt.Fprintf(stderr, "csvstat: %v\n", err)
		return exitFailure
	}
	return exitOK
}

// parseArgs parses the command line with a local FlagSet and checks the usage rules. On a usage
// error it writes the reason and both synopses to stderr and reports false.
func parseArgs(args []string, stderr io.Writer) (request, bool) {
	fs := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	column := fs.String("column", "", "")
	freqName := fs.String("freq", "", "")
	top := fs.Int("top", defaultTop, "")

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			err = errors.New("no help option is defined")
		}
		usageError(stderr, err.Error())
		return request{}, false
	}

	set := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { set[f.Name] = true })

	switch {
	case fs.NArg() != 1:
		usageError(stderr, "expected exactly one file argument")
	case set["column"] && set["freq"]:
		usageError(stderr, "--column and --freq cannot be combined")
	case set["top"] && !set["freq"]:
		usageError(stderr, "--top requires --freq")
	case set["top"] && *top < 1:
		usageError(stderr, "--top must be a positive integer")
	default:
		return request{
			file:      fs.Arg(0),
			column:    *column,
			hasColumn: set["column"],
			freqName:  *freqName,
			hasFreq:   set["freq"],
			top:       *top,
		}, true
	}
	return request{}, false
}

func usageError(stderr io.Writer, reason string) {
	fmt.Fprintf(stderr, "csvstat: %s\n", reason)
	fmt.Fprintf(stderr, "usage: %s\n", profileSynopsis)
	fmt.Fprintf(stderr, "       %s\n", freqSynopsis)
}

func absentColumn(stderr io.Writer, file, name string) int {
	fmt.Fprintf(stderr, "csvstat: %s: no column named %q\n", file, name)
	return exitRejected
}

// profiledColumns returns the header indexes to profile: every column, or only the one named by
// --column. It reports false when the named column is absent.
func profiledColumns(t reader.Table, req request) ([]int, bool) {
	if !req.hasColumn {
		all := make([]int, len(t.Header))
		for i := range all {
			all[i] = i
		}
		return all, true
	}
	j, found := columnIndex(t.Header, req.column)
	return []int{j}, found
}

// columnIndex finds a header name by exact, case-sensitive comparison.
func columnIndex(header []string, name string) (int, bool) {
	for i, h := range header {
		if h == name {
			return i, true
		}
	}
	return 0, false
}

func columnValues(t reader.Table, j int) []string {
	values := make([]string, len(t.Rows))
	for i, row := range t.Rows {
		values[i] = row[j]
	}
	return values
}
