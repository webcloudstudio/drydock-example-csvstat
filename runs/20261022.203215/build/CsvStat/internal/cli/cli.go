// Package cli owns command-line parsing and process-facing diagnostics.
package cli

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type command struct {
	file   string
	column string
	freq   string
	top    int
}

const usage = `Usage:
  csvstat [--column <name>] <file>
  csvstat --freq <name> [--top <n>] <file>

Options:
  --column <name>  profile one column
  --freq <name>    list the most frequent values in a column
  --top <n>        limit a frequency report to a positive number of values
  --help           show this help

Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
`

// Run parses args and returns the process exit code for csvstat.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 1 && args[0] == "--help" {
		writeUsage(stderr)
		return 0
	}

	cmd, err := parse(args)
	if err != nil {
		writeDiagnostic(stderr, err)
		return 2
	}

	header, err := readHeader(cmd.file)
	if err != nil {
		writeDiagnostic(stderr, err)
		if _, ok := err.(csvInputError); ok {
			return 3
		}
		return 1
	}
	if len(header) == 0 {
		writeDiagnostic(stderr, fmt.Errorf("CSV has no header"))
		return 3
	}
	if cmd.column != "" && !hasColumn(header, cmd.column) {
		writeDiagnostic(stderr, fmt.Errorf("column %q is absent from the header", cmd.column))
		return 3
	}
	if cmd.freq != "" && !hasColumn(header, cmd.freq) {
		writeDiagnostic(stderr, fmt.Errorf("column %q is absent from the header", cmd.freq))
		return 3
	}

	if cmd.freq != "" {
		_, _ = fmt.Fprintf(stdout, "frequency: %s\n", cmd.freq)
	} else if cmd.column != "" {
		_, _ = fmt.Fprintf(stdout, "profile: %s\n", cmd.column)
	} else {
		_, _ = fmt.Fprintln(stdout, "profile: all columns")
	}
	return 0
}

func parse(args []string) (command, error) {
	var cmd command
	positional := false
	topSet := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if positional {
			return command{}, fmt.Errorf("options must precede the file argument")
		}
		switch arg {
		case "--column", "--freq":
			if i+1 >= len(args) || args[i+1] == "" || strings.HasPrefix(args[i+1], "-") {
				return command{}, fmt.Errorf("%s requires a value", arg)
			}
			value := args[i+1]
			i++
			if arg == "--column" {
				if cmd.column != "" {
					return command{}, fmt.Errorf("--column may be specified only once")
				}
				cmd.column = value
			} else {
				if cmd.freq != "" {
					return command{}, fmt.Errorf("--freq may be specified only once")
				}
				cmd.freq = value
			}
		case "--top":
			if topSet {
				return command{}, fmt.Errorf("--top may be specified only once")
			}
			if i+1 >= len(args) || args[i+1] == "" {
				return command{}, fmt.Errorf("--top requires a positive integer")
			}
			top, err := strconv.Atoi(args[i+1])
			if err != nil || top <= 0 {
				return command{}, fmt.Errorf("--top requires a positive integer")
			}
			cmd.top = top
			topSet = true
			i++
		case "--help", "-h":
			return command{}, fmt.Errorf("help cannot be combined with other arguments")
		default:
			if strings.HasPrefix(arg, "-") {
				return command{}, fmt.Errorf("unknown option %q", arg)
			}
			if cmd.file != "" {
				return command{}, fmt.Errorf("exactly one file argument is required")
			}
			cmd.file = arg
			positional = true
		}
	}
	if cmd.file == "" {
		return command{}, fmt.Errorf("exactly one file argument is required")
	}
	if cmd.column != "" && cmd.freq != "" {
		return command{}, fmt.Errorf("--column and --freq are mutually exclusive")
	}
	if cmd.top > 0 && cmd.freq == "" {
		return command{}, fmt.Errorf("--top is valid only with --freq")
	}
	return cmd, nil
}

func readHeader(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, csvInputError{err: err}
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records[0], nil
}

type csvInputError struct{ err error }

func (e csvInputError) Error() string { return fmt.Sprintf("invalid CSV: %v", e.err) }

func hasColumn(header []string, wanted string) bool {
	for _, name := range header {
		if name == wanted {
			return true
		}
	}
	return false
}

func writeUsage(stderr io.Writer) {
	_, _ = fmt.Fprint(stderr, usage)
}

func writeDiagnostic(stderr io.Writer, err error) {
	_, _ = fmt.Fprintf(stderr, "csvstat: %v\n%s", err, usage)
}
