package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"
)

// Options is the parsed command line. ColumnSet, FreqSet and TopSet record
// whether the option was given, so an empty value is distinguishable from an
// absent option.
type Options struct {
	Column    string
	ColumnSet bool
	Freq      string
	FreqSet   bool
	Top       int
	TopSet    bool
	File      string
}

var errRepeated = errors.New("option given more than once")

// stringOption is a flag value that records presence and refuses repetition.
type stringOption struct {
	value *string
	set   *bool
}

func (o *stringOption) String() string {
	if o == nil || o.value == nil {
		return ""
	}
	return *o.value
}

func (o *stringOption) Set(v string) error {
	if *o.set {
		return errRepeated
	}
	*o.value = v
	*o.set = true
	return nil
}

// intOption is a flag value holding a decimal integer that refuses repetition.
type intOption struct {
	value *int
	set   *bool
}

func (o *intOption) String() string {
	if o == nil || o.value == nil {
		return ""
	}
	return strconv.Itoa(*o.value)
}

func (o *intOption) Set(v string) error {
	if *o.set {
		return errRepeated
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("%q is not a decimal integer", v)
	}
	*o.value = n
	*o.set = true
	return nil
}

// ParseArgs parses options followed by exactly one file argument. Parsing stops
// at the first non-option argument, which must be the file, so any later
// argument, including an option, is an error.
func ParseArgs(args []string) (Options, error) {
	var opts Options
	fs := flag.NewFlagSet("csvstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(&stringOption{&opts.Column, &opts.ColumnSet}, "column", "")
	fs.Var(&stringOption{&opts.Freq, &opts.FreqSet}, "freq", "")
	fs.Var(&intOption{&opts.Top, &opts.TopSet}, "top", "")
	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}

	rest := fs.Args()
	switch len(rest) {
	case 0:
		return Options{}, errors.New("missing file argument")
	case 1:
		opts.File = rest[0]
	default:
		return Options{}, fmt.Errorf("unexpected argument %q after file or extra file", rest[1])
	}
	if opts.ColumnSet && opts.FreqSet {
		return Options{}, errors.New("--column cannot be combined with --freq")
	}
	if opts.TopSet && !opts.FreqSet {
		return Options{}, errors.New("--top requires --freq")
	}
	if opts.TopSet && opts.Top <= 0 {
		return Options{}, fmt.Errorf("--top %d is not a positive integer", opts.Top)
	}
	return opts, nil
}
