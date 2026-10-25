package cli

import "testing"

func TestParseArgsFileOnly(t *testing.T) {
	opts, err := ParseArgs([]string{"data.csv"})
	if err != nil {
		t.Fatalf("ParseArgs: unexpected error: %v", err)
	}
	if want := (Options{File: "data.csv"}); opts != want {
		t.Fatalf("Options = %+v, want %+v", opts, want)
	}
}

func TestParseArgsColumn(t *testing.T) {
	opts, err := ParseArgs([]string{"--column", "name", "data.csv"})
	if err != nil {
		t.Fatalf("ParseArgs: unexpected error: %v", err)
	}
	if want := (Options{Column: "name", ColumnSet: true, File: "data.csv"}); opts != want {
		t.Fatalf("Options = %+v, want %+v", opts, want)
	}
}

func TestParseArgsFreqTop(t *testing.T) {
	opts, err := ParseArgs([]string{"--freq", "name", "--top", "3", "data.csv"})
	if err != nil {
		t.Fatalf("ParseArgs: unexpected error: %v", err)
	}
	want := Options{Freq: "name", FreqSet: true, Top: 3, TopSet: true, File: "data.csv"}
	if opts != want {
		t.Fatalf("Options = %+v, want %+v", opts, want)
	}
}

func TestParseArgsEmptyValueIsSet(t *testing.T) {
	opts, err := ParseArgs([]string{"--column", "", "data.csv"})
	if err != nil {
		t.Fatalf("ParseArgs: unexpected error: %v", err)
	}
	if !opts.ColumnSet || opts.Column != "" {
		t.Fatalf("Options = %+v, want ColumnSet with empty Column", opts)
	}
}

func TestParseArgsMissingFile(t *testing.T) {
	if _, err := ParseArgs([]string{"--column", "name"}); err == nil {
		t.Fatal("ParseArgs with no file: got nil error, want error")
	}
}

func TestParseArgsExtraFile(t *testing.T) {
	if _, err := ParseArgs([]string{"a.csv", "b.csv"}); err == nil {
		t.Fatal("ParseArgs with two files: got nil error, want error")
	}
}

func TestParseArgsUnknownOption(t *testing.T) {
	if _, err := ParseArgs([]string{"--frobnicate", "data.csv"}); err == nil {
		t.Fatal("ParseArgs with unknown option: got nil error, want error")
	}
}

func TestParseArgsMissingValue(t *testing.T) {
	for _, args := range [][]string{{"--top"}, {"--freq"}, {"--top", "data.csv"}} {
		if _, err := ParseArgs(args); err == nil {
			t.Errorf("ParseArgs(%q): got nil error, want error", args)
		}
	}
}

func TestParseArgsOptionsAfterFile(t *testing.T) {
	if _, err := ParseArgs([]string{"data.csv", "--column", "name"}); err == nil {
		t.Fatal("ParseArgs with option after file: got nil error, want error")
	}
}

func TestParseArgsRepeatedOption(t *testing.T) {
	if _, err := ParseArgs([]string{"--column", "a", "--column", "b", "data.csv"}); err == nil {
		t.Fatal("ParseArgs with repeated --column: got nil error, want error")
	}
}

func TestParseArgsColumnWithFreqIsError(t *testing.T) {
	if _, err := ParseArgs([]string{"--column", "a", "--freq", "a", "data.csv"}); err == nil {
		t.Fatal("ParseArgs with --column and --freq: got nil error, want error")
	}
}

func TestParseArgsTopWithoutFreqIsError(t *testing.T) {
	if _, err := ParseArgs([]string{"--top", "3", "data.csv"}); err == nil {
		t.Fatal("ParseArgs with --top and no --freq: got nil error, want error")
	}
}

func TestParseArgsTopMustBePositive(t *testing.T) {
	for _, v := range []string{"0", "-1"} {
		if _, err := ParseArgs([]string{"--freq", "a", "--top", v, "data.csv"}); err == nil {
			t.Errorf("ParseArgs with --top %s: got nil error, want error", v)
		}
	}
}

func TestParseArgsTopMustBeDecimal(t *testing.T) {
	if _, err := ParseArgs([]string{"--top", "0x10", "data.csv"}); err == nil {
		t.Fatal("ParseArgs with --top 0x10: got nil error, want error")
	}
}
