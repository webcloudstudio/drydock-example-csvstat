package cli

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"csvstat/internal/csvstat"
)

const (
	ExitOK          = 0
	ExitOperational = 1
	ExitUsage       = 2
	ExitRejected    = 3
)

func Run(args []string, stdout, stderr io.Writer) int {
	f := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	f.SetOutput(stderr)
	f.Usage = func() { _, _ = io.WriteString(stderr, usage()) }
	help := f.Bool("help", false, "show command usage")
	column := f.String("column", "", "profile one column by name")
	freq := f.String("freq", "", "show frequent values for one column")
	topText := f.String("top", "5", "maximum number of frequent values to show (default 5)")
	if err := f.Parse(args); err != nil {
		return ExitUsage
	}
	if *help {
		if _, err := io.WriteString(stdout, usage()); err != nil {
			return ExitOperational
		}
		return ExitOK
	}
	if f.NArg() != 1 || (*column != "" && *freq != "") {
		_, _ = io.WriteString(stderr, usage())
		return ExitUsage
	}
	top := 5
	if hasFlag(args, "top") {
		var err error
		top, err = strconv.Atoi(*topText)
		if err != nil || top <= 0 {
			fmt.Fprintln(stderr, "csvstat: --top must be a positive integer")
			return ExitUsage
		}
	}
	if !hasFlag(args, "freq") && hasFlag(args, "top") {
		fmt.Fprintln(stderr, "csvstat: --top requires --freq")
		return ExitUsage
	}
	if hasFlag(args, "freq") && *freq == "" {
		fmt.Fprintln(stderr, "csvstat: --freq requires a column name")
		return ExitUsage
	}
	path := f.Arg(0)
	file, err := os.Open(path)
	if err != nil {
		diagnostic(stderr, path, err)
		return ExitOperational
	}
	defer file.Close()
	in, err := csvstat.Load(file)
	if err != nil {
		diagnostic(stderr, path, err)
		return ExitRejected
	}
	var output bytes.Buffer
	if *freq != "" {
		v, e := csvstat.Frequencies(in, *freq, top)
		if e != nil {
			diagnostic(stderr, path, e)
			return ExitRejected
		}
		err = csvstat.RenderFrequency(&output, v)
	} else {
		r, e := csvstat.Profile(in, *column)
		if e != nil {
			diagnostic(stderr, path, e)
			return ExitRejected
		}
		err = csvstat.RenderProfile(&output, r)
	}
	if err != nil {
		diagnostic(stderr, path, err)
		return ExitOperational
	}
	if _, err = io.Copy(stdout, &output); err != nil {
		diagnostic(stderr, path, err)
		return ExitOperational
	}
	return ExitOK
}
func diagnostic(w io.Writer, path string, err error) { fmt.Fprintf(w, "csvstat: %s: %v\n", path, err) }
func hasFlag(args []string, name string) bool {
	for _, a := range args {
		if a == "--" {
			return false
		}
		if a == "--"+name || len(a) > len(name)+3 && a[:len(name)+3] == "--"+name+"=" {
			return true
		}
	}
	return false
}
func usage() string {
	return "Usage:\n  csvstat [--column <name>] <file>\n  csvstat --freq <name> [--top <n>] <file>\n\nOptions:\n  --column <name>  profile one column by name\n  --freq <name>    show frequent values for one column\n  --top <n>        maximum number of frequent values to show (default 5)\n"
}
