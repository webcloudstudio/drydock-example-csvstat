package cli

import (
	"reflect"
	"testing"
)

func TestParseHelp(t *testing.T) {
	req, err := parse([]string{"--help"})
	if err != nil {
		t.Fatalf("parse(--help) error = %v", err)
	}
	if !req.Help {
		t.Errorf("Help = false, want true")
	}
}

func TestParseColumnForm(t *testing.T) {
	req, err := parse([]string{"--column", "qty", "orders.csv"})
	if err != nil {
		t.Fatal(err)
	}
	want := Request{Mode: ModeProfile, Column: "qty", Path: "orders.csv", Top: defaultTop}
	if !reflect.DeepEqual(req, want) {
		t.Errorf("parse() = %+v, want %+v", req, want)
	}
}

func TestParseFreqForm(t *testing.T) {
	req, err := parse([]string{"--freq", "sensor", "sensors.csv"})
	if err != nil {
		t.Fatal(err)
	}
	want := Request{Mode: ModeFrequency, Column: "sensor", Path: "sensors.csv", Top: 5}
	if !reflect.DeepEqual(req, want) {
		t.Errorf("parse() = %+v, want %+v", req, want)
	}
}

func TestParseProfileAllColumns(t *testing.T) {
	req, err := parse([]string{"orders.csv"})
	if err != nil {
		t.Fatal(err)
	}
	if !req.AllColumns || req.Mode != ModeProfile || req.Path != "orders.csv" {
		t.Errorf("parse() = %+v, want profile of all columns", req)
	}
}

func TestParseTopValues(t *testing.T) {
	cases := map[string]int{
		"1":                       1,
		"3":                       3,
		"007":                     7,
		"99999999999999999999999": int(^uint(0) >> 1), // too large for int: unbounded
	}
	for in, want := range cases {
		req, err := parse([]string{"--freq", "c", "--top", in, "f.csv"})
		if err != nil {
			t.Errorf("--top %q: error = %v", in, err)
			continue
		}
		if req.Top != want {
			t.Errorf("--top %q: Top = %d, want %d", in, req.Top, want)
		}
	}
}

func TestParseUsageErrors(t *testing.T) {
	cases := map[string][]string{
		"no file":              {},
		"two files":            {"a.csv", "b.csv"},
		"unknown option":       {"--frobnicate", "a.csv"},
		"column missing value": {"a.csv", "--column"},
		"freq missing value":   {"a.csv", "--freq"},
		"top missing value":    {"--freq", "q", "a.csv", "--top"},
		"top zero":             {"--freq", "q", "--top", "0", "a.csv"},
		"top negative":         {"--freq", "q", "--top", "-4", "a.csv"},
		"top text":             {"--freq", "q", "--top", "many", "a.csv"},
		"top fraction":         {"--freq", "q", "--top", "2.5", "a.csv"},
		"top empty":            {"--freq", "q", "--top", "", "a.csv"},
		"top without freq":     {"--top", "3", "a.csv"},
		"column with freq":     {"--column", "q", "--freq", "q", "a.csv"},
	}
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := parse(args); err == nil {
				t.Errorf("parse(%q) error = nil, want usage error", args)
			}
		})
	}
}
