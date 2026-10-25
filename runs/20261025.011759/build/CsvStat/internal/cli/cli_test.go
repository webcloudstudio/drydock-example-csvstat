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

const profileHeaderLine = "column,type,rows,nulls,distinct,min,max,sum,mean,median"

// employees is a small fixture with a header, typed columns, nulls, and a quoted comma.
const employees = "  id , name ,hired,salary,remote,notes\n" +
	"3,Ana,2025-01-05,1.50,true,\"likes tea, coffee\"\n" +
	"1,Ben,2024-12-31,-2.25,FALSE,NA\n" +
	"2,  Cy  ,2025-03-01,10.00,true,\n" +
	"10,Dee,NULL,N/A,true,plain\n"

func writeCSV(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func runCLI(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	var out, errOut bytes.Buffer
	code = Run(args, &out, &errOut)
	return code, out.String(), errOut.String()
}

// parseOutput reads stdout back through encoding/csv, so checks compare parsed records.
func parseOutput(t *testing.T, stdout string) [][]string {
	t.Helper()
	records, err := csv.NewReader(strings.NewReader(stdout)).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, stdout)
	}
	return records
}

func TestProfileHeader(t *testing.T) {
	code, stdout, _ := runCLI(t, writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records := parseOutput(t, stdout)
	want := strings.Split(profileHeaderLine, ",")
	if !reflect.DeepEqual(records[0], want) {
		t.Errorf("header = %q, want %q", records[0], want)
	}
}

func TestProfileRowOrder(t *testing.T) {
	code, stdout, _ := runCLI(t, writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records := parseOutput(t, stdout)
	var names []string
	for _, r := range records[1:] {
		names = append(names, r[0])
	}
	if want := []string{"id", "name", "hired", "salary", "remote", "notes"}; !reflect.DeepEqual(names, want) {
		t.Errorf("column order = %q, want %q", names, want)
	}
}

func TestProfileQuoting(t *testing.T) {
	code, stdout, _ := runCLI(t, writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records := parseOutput(t, stdout)
	notes := records[len(records)-1]
	if notes[0] != "notes" {
		t.Fatalf("last row = %q, want the notes column", notes)
	}
	// The quoted comma value reads back as one field: notes has 4 rows, 2 nulls, 2 distinct values.
	if notes[2] != "4" || notes[3] != "2" || notes[4] != "2" {
		t.Errorf("notes counts = rows %q nulls %q distinct %q, want 4, 2, 2", notes[2], notes[3], notes[4])
	}
	if notes[5] != "likes tea, coffee" || notes[6] != "plain" {
		t.Errorf("notes min,max = %q,%q, want %q,plain", notes[5], notes[6], "likes tea, coffee")
	}

	code, stdout, _ = runCLI(t, "--column", "notes", writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records = parseOutput(t, stdout)
	if len(records) != 2 {
		t.Fatalf("rows = %d, want header plus one", len(records))
	}
	if records[1][0] != "notes" {
		t.Errorf("column row = %q, want notes", records[1])
	}
}

func TestProfileFieldsOfNumericColumn(t *testing.T) {
	code, stdout, _ := runCLI(t, writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	var salary []string
	for _, r := range parseOutput(t, stdout) {
		if r[0] == "salary" {
			salary = r
		}
	}
	// N/A is null; values 1.50, -2.25, 10.00 give sum 9.25 and mean 3.08 (3.0833 rounded).
	want := []string{"salary", "number", "4", "1", "3", "-2.25", "10.00", "9.25", "3.08", "1.50"}
	if !reflect.DeepEqual(salary, want) {
		t.Errorf("salary row = %q, want %q", salary, want)
	}
}

func TestColumnSelectMatchesFullProfile(t *testing.T) {
	path := writeCSV(t, employees)
	_, full, _ := runCLI(t, path)
	records := parseOutput(t, full)
	byName := make(map[string][]string)
	for _, r := range records[1:] {
		byName[r[0]] = r
	}
	for name, row := range byName {
		code, stdout, _ := runCLI(t, "--column", name, path)
		if code != 0 {
			t.Fatalf("--column %s exit = %d, want 0", name, code)
		}
		got := parseOutput(t, stdout)
		if len(got) != 2 || !reflect.DeepEqual(got[1], row) {
			t.Errorf("--column %s = %q, want %q", name, got, row)
		}
	}
}

func TestColumnSelectTrimmedHeaderName(t *testing.T) {
	code, stdout, _ := runCLI(t, "--column", "name", writeCSV(t, employees))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records := parseOutput(t, stdout)
	if records[1][0] != "name" {
		t.Errorf("row = %q, want the name column", records[1])
	}
}

func TestColumnAbsentExitsThree(t *testing.T) {
	code, stdout, _ := runCLI(t, "--column", "nosuch", writeCSV(t, employees))
	if code != 3 || stdout != "" {
		t.Errorf("exit = %d, stdout = %q; want 3 and nothing on stdout", code, stdout)
	}
}

func TestFrequencyOrderAndNulls(t *testing.T) {
	path := writeCSV(t, "v\nb\na\nB\nc\na\nb\nNA\n")
	code, stdout, _ := runCLI(t, "--freq", "v", path)
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	records := parseOutput(t, stdout)
	want := [][]string{{"value", "count"}, {"a", "2"}, {"b", "2"}, {"B", "1"}, {"c", "1"}}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("frequency rows = %q, want %q", records, want)
	}
}

func TestFrequencyDefaultLimitIsFive(t *testing.T) {
	path := writeCSV(t, "v\na\nb\nc\nd\ne\nf\ng\n")
	code, stdout, _ := runCLI(t, "--freq", "v", path)
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	if got := len(parseOutput(t, stdout)) - 1; got != 5 {
		t.Errorf("rows = %d, want 5", got)
	}
}

func TestFrequencyTopLimitsRows(t *testing.T) {
	path := writeCSV(t, "v\na\nb\nb\nc\nc\nc\n")
	code, stdout, _ := runCLI(t, "--top", "2", "--freq", "v", path)
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	want := [][]string{{"value", "count"}, {"c", "3"}, {"b", "2"}}
	if got := parseOutput(t, stdout); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestFrequencyAllNullWritesHeaderOnly(t *testing.T) {
	code, stdout, _ := runCLI(t, "--freq", "v", writeCSV(t, "v\nNA\n\nnull\n"))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	if want := [][]string{{"value", "count"}}; !reflect.DeepEqual(parseOutput(t, stdout), want) {
		t.Errorf("rows = %q, want header only", parseOutput(t, stdout))
	}
}

func TestFrequencyAbsentColumnExitsThree(t *testing.T) {
	code, stdout, _ := runCLI(t, "--freq", "nosuch", writeCSV(t, employees))
	if code != 3 || stdout != "" {
		t.Errorf("exit = %d, stdout = %q; want 3 and nothing on stdout", code, stdout)
	}
}

func TestHeaderOnlyProfilesEveryColumnAsEmpty(t *testing.T) {
	code, stdout, _ := runCLI(t, writeCSV(t, "sku,price\n"))
	if code != 0 {
		t.Fatalf("exit = %d, want 0", code)
	}
	want := [][]string{
		strings.Split(profileHeaderLine, ","),
		{"sku", "empty", "0", "0", "0", "", "", "", "", ""},
		{"price", "empty", "0", "0", "0", "", "", "", "", ""},
	}
	if got := parseOutput(t, stdout); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestUsageErrorsExitTwoWithOneLine(t *testing.T) {
	path := writeCSV(t, "a,b\n1,2\n")
	cases := map[string][]string{
		"no file":              {},
		"two files":            {path, path},
		"unknown option":       {"--frobnicate", path},
		"help option":          {"-h", path},
		"option without value": {"--column"},
		"freq without value":   {"--freq"},
		"column with freq":     {"--column", "a", "--freq", "a", path},
		"top without freq":     {"--top", "3", path},
		"top zero":             {"--freq", "a", "--top", "0", path},
		"top negative":         {"--freq", "a", "--top", "-1", path},
		"top not a number":     {"--freq", "a", "--top", "abc", path},
		"top decimal":          {"--freq", "a", "--top", "2.5", path},
		"top empty":            {"--freq", "a", "--top", "", path},
		"option after file":    {path, "--column", "a"},
	}
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			code, stdout, stderr := runCLI(t, args...)
			if code != 2 || stdout != "" {
				t.Errorf("exit = %d, stdout = %q; want 2 and nothing on stdout", code, stdout)
			}
			if lines := nonEmptyLines(stderr); len(lines) != 1 {
				t.Errorf("stderr = %q, want exactly one line", stderr)
			}
		})
	}
}

func TestReadFailureExitsOneWithNothingOnStdout(t *testing.T) {
	cases := map[string]string{
		"missing file": filepath.Join(t.TempDir(), "absent.csv"),
		"directory":    t.TempDir(),
	}
	for name, path := range cases {
		t.Run(name, func(t *testing.T) {
			code, stdout, _ := runCLI(t, path)
			if code != 1 || stdout != "" {
				t.Errorf("exit = %d, stdout = %q; want 1 and nothing on stdout", code, stdout)
			}
		})
	}
}

func TestRejectedInputExitsThreeWithNothingOnStdout(t *testing.T) {
	cases := map[string]string{
		"empty file":   "",
		"ragged":       "a,b\n1,2\n3,4,5\n",
		"empty header": "a,,c\n1,2,3\n",
		"repeated":     "a, a ,c\n1,2,3\n",
	}
	for name, content := range cases {
		t.Run(name, func(t *testing.T) {
			code, stdout, stderr := runCLI(t, writeCSV(t, content))
			if code != 3 || stdout != "" {
				t.Errorf("exit = %d, stdout = %q; want 3 and nothing on stdout", code, stdout)
			}
			if len(nonEmptyLines(stderr)) == 0 {
				t.Errorf("stderr is empty, want a diagnostic")
			}
		})
	}
}

func nonEmptyLines(s string) []string {
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
