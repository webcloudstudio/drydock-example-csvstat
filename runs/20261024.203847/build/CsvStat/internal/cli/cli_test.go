package cli

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// writeFile writes content to a file in a fresh temporary directory and returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

// runCLI runs the command and returns its exit code, stdout, and stderr.
func runCLI(args ...string) (int, string, string) {
	var stdout, stderr bytes.Buffer
	code := Run(args, &stdout, &stderr)
	return code, stdout.String(), stderr.String()
}

func parseCSV(t *testing.T, s string) [][]string {
	t.Helper()
	recs, err := csv.NewReader(strings.NewReader(s)).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%q", err, s)
	}
	return recs
}

func TestRunReturnsExitCode(t *testing.T) {
	good := writeFile(t, "a,b\n1,2\n")
	ragged := writeFile(t, "a,b\n1,2\n3\n")
	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "no arguments", args: nil, want: 2},
		{name: "two files", args: []string{good, good}, want: 2},
		{name: "unknown flag", args: []string{"--bogus", good}, want: 2},
		{name: "help flag", args: []string{"-h", good}, want: 2},
		{name: "long help flag", args: []string{"--help", good}, want: 2},
		{name: "one file", args: []string{good}, want: 0},
		{name: "missing file", args: []string{filepath.Join(t.TempDir(), "none.csv")}, want: 1},
		{name: "directory path", args: []string{t.TempDir()}, want: 1},
		{name: "ragged record", args: []string{ragged}, want: 3},
		{name: "absent column", args: []string{"--column", "zz", good}, want: 3},
		{name: "absent freq column", args: []string{"--freq", "zz", good}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _, _ := runCLI(tt.args...); got != tt.want {
				t.Errorf("Run(%q) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func TestFailuresWriteNothingToStdout(t *testing.T) {
	ragged := writeFile(t, "a,b\n1,2\n3,4,5\n")
	tests := [][]string{
		nil,
		{"--top", "3", ragged},
		{"--column", "zz", ragged},
		{"--freq", "a", ragged},
		{filepath.Join(t.TempDir(), "none.csv")},
	}
	for _, args := range tests {
		code, stdout, _ := runCLI(args...)
		if code == 0 {
			t.Fatalf("Run(%q) = 0, want a failure", args)
		}
		if stdout != "" {
			t.Errorf("Run(%q) wrote %q to stdout, want nothing", args, stdout)
		}
	}
}

func TestUsageErrorsNameBothSynopses(t *testing.T) {
	good := writeFile(t, "a\n1\n")
	_, _, stderr := runCLI("--frobnicate", good)
	for _, s := range []string{"csvstat [--column <name>] <file>", "csvstat --freq <name> [--top <n>] <file>"} {
		if !strings.Contains(stderr, s) {
			t.Errorf("stderr %q lacks synopsis %q", stderr, s)
		}
	}
}

func TestUsageErrorsDecidedBeforeFileOpened(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "none.csv")
	if code, _, _ := runCLI("--top", "3", missing); code != 2 {
		t.Errorf("exit = %d, want 2", code)
	}
}

func TestUsageRules(t *testing.T) {
	good := writeFile(t, "qty\n1\n")
	tests := []struct {
		name string
		args []string
	}{
		{"missing option value column", []string{"--column"}},
		{"missing option value freq", []string{"--freq"}},
		{"missing option value top", []string{"--freq", "qty", good, "--top"}},
		{"top zero", []string{"--freq", "qty", "--top", "0", good}},
		{"top negative", []string{"--freq", "qty", "--top", "-1", good}},
		{"top not a number", []string{"--freq", "qty", "--top", "abc", good}},
		{"top fractional", []string{"--freq", "qty", "--top", "1.5", good}},
		{"top without freq", []string{"--top", "3", good}},
		{"column with freq", []string{"--column", "qty", "--freq", "qty", good}},
		{"file before option", []string{good, "--column", "qty"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if code, _, _ := runCLI(tt.args...); code != 2 {
				t.Errorf("Run(%q) = %d, want 2", tt.args, code)
			}
		})
	}
}

func TestProfileAllColumnsInFileOrder(t *testing.T) {
	path := writeFile(t, "zeta,alpha,mid\n1,x,true\n2,y,false\n")
	code, stdout, _ := runCLI(path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, stdout)
	if !reflect.DeepEqual(rows[0], []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}) {
		t.Errorf("header = %q", rows[0])
	}
	var names []string
	for _, r := range rows[1:] {
		names = append(names, r[0])
	}
	if !reflect.DeepEqual(names, []string{"zeta", "alpha", "mid"}) {
		t.Errorf("columns = %q", names)
	}
}

func TestProfileValuesRoundTrip(t *testing.T) {
	path := writeFile(t, "id,qty,note\n1,10,a\n2,20,\n3,NA,b\n")
	code, stdout, _ := runCLI(path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, stdout)
	want := map[string][]string{
		"id":   {"integer", "3", "0", "3", "1", "3", "6", "2.00", "2.00"},
		"qty":  {"integer", "3", "1", "2", "10", "20", "30", "15.00", "15.00"},
		"note": {"text", "3", "1", "2", "a", "b", "", "", ""},
	}
	for _, r := range rows[1:] {
		if !reflect.DeepEqual(r[1:], want[r[0]]) {
			t.Errorf("row %q = %q, want %q", r[0], r[1:], want[r[0]])
		}
	}
}

func TestColumnProfilesOnlyNamedColumn(t *testing.T) {
	path := writeFile(t, "id,qty\n1,10\n2,20\n")
	code, stdout, _ := runCLI("--column", "qty", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, stdout)
	if len(rows) != 2 || rows[1][0] != "qty" {
		t.Errorf("rows = %q, want header plus qty only", rows)
	}
}

func TestColumnMatchesTrimmedHeader(t *testing.T) {
	path := writeFile(t, "id,  qty  \n1,10\n")
	code, stdout, _ := runCLI("--column", "qty", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if rows := parseCSV(t, stdout); rows[1][0] != "qty" {
		t.Errorf("rows = %q", rows)
	}
}

func TestColumnNameIsCaseSensitive(t *testing.T) {
	path := writeFile(t, "qty\n1\n")
	if code, _, _ := runCLI("--column", "QTY", path); code != 3 {
		t.Errorf("exit = %d, want 3", code)
	}
}

func TestHeaderOnlyFileProfilesEmptyColumns(t *testing.T) {
	path := writeFile(t, "sku,price,in_stock\n")
	code, stdout, _ := runCLI(path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, stdout)
	if len(rows) != 4 || rows[1][1] != "empty" || rows[1][2] != "0" {
		t.Errorf("rows = %q", rows)
	}
}

func TestFreqListsMostFrequentDefaultFive(t *testing.T) {
	var b strings.Builder
	b.WriteString("v\n")
	for i := 0; i < 8; i++ {
		b.WriteString("item" + string(rune('a'+i)) + "\n")
	}
	path := writeFile(t, b.String())
	code, stdout, _ := runCLI("--freq", "v", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, stdout)
	if len(rows)-1 != 5 || rows[0][0] != "value" {
		t.Errorf("rows = %q, want header plus five values", rows)
	}
}

func TestFreqOrdersByCountThenValue(t *testing.T) {
	path := writeFile(t, "k\nb\na\nb\nc\nc\nd\n")
	code, stdout, _ := runCLI("--freq", "k", "--top", "3", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	want := [][]string{{"value", "count"}, {"b", "2"}, {"c", "2"}, {"a", "1"}}
	if got := parseCSV(t, stdout); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestFreqSkipsNullsAndTrims(t *testing.T) {
	path := writeFile(t, "customer\n  Globex  \nGlobex\n\nNA\nNULL\nOther\n")
	code, stdout, _ := runCLI("--freq", "customer", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	want := [][]string{{"value", "count"}, {"Globex", "2"}, {"Other", "1"}}
	if got := parseCSV(t, stdout); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestFreqAllNullColumnWritesHeaderOnly(t *testing.T) {
	path := writeFile(t, "k,other\n,x\nNULL,y\n")
	code, stdout, _ := runCLI("--freq", "k", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if stdout != "value,count\n" {
		t.Errorf("stdout = %q, want header only", stdout)
	}
}

func TestFreqTopLargerThanDistinct(t *testing.T) {
	path := writeFile(t, "v\na\nb\n")
	code, stdout, _ := runCLI("--freq", "v", "--top", "50", path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if rows := parseCSV(t, stdout); len(rows)-1 != 2 {
		t.Errorf("rows = %q, want two values", rows)
	}
}

func TestRaggedRecordDiagnosticNamesLine(t *testing.T) {
	path := writeFile(t, "sku,price\nA1,9.99\nA2,4.50,extra\n")
	code, _, stderr := runCLI(path)
	if code != 3 {
		t.Fatalf("exit = %d, want 3", code)
	}
	if !strings.Contains(stderr, "line 3") || !strings.Contains(stderr, filepath.Base(path)) {
		t.Errorf("stderr = %q, want the file and line 3", stderr)
	}
}

func TestUnreadableFileDiagnosticNamesPath(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "none.csv")
	code, _, stderr := runCLI(missing)
	if code != 1 || !strings.Contains(stderr, missing) {
		t.Errorf("exit = %d, stderr = %q", code, stderr)
	}
}

func TestInvalidUTF8IsAccepted(t *testing.T) {
	path := writeFile(t, "name\n\xff\xfe\n")
	if code, _, _ := runCLI(path); code != 0 {
		t.Errorf("exit = %d, want 0", code)
	}
}

func TestNumericsAreExactForLargeIntegers(t *testing.T) {
	path := writeFile(t, "n\n9223372036854775807\n9223372036854775807\n9223372036854775807\n")
	code, stdout, _ := runCLI(path)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if rows := parseCSV(t, stdout); rows[1][7] != "27670116110564327421" {
		t.Errorf("sum = %q", rows[1][7])
	}
}
