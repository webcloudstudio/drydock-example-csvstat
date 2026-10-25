package profile

import (
	"reflect"
	"testing"

	"csvstat/internal/csvdata"
)

// rowsOf builds a table with one column per name and returns its profile rows.
func rowsOf(t *testing.T, names []string, columns ...[]string) [][]string {
	t.Helper()
	table := csvdata.Table{Header: names}
	for i := range columns[0] {
		var fields []string
		for _, col := range columns {
			fields = append(fields, col[i])
		}
		table.Records = append(table.Records, csvdata.Record{Line: i + 2, Fields: fields})
	}
	return Rows(table)
}

// rowOf returns the profile row for a single column of values.
func rowOf(t *testing.T, values ...string) []string {
	t.Helper()
	return rowsOf(t, []string{"c"}, values)[0]
}

func TestHeaderIsFixedTenFields(t *testing.T) {
	want := []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}
	if !reflect.DeepEqual(Header, want) {
		t.Errorf("Header = %q, want %q", Header, want)
	}
}

func TestRowsFollowFileOrderAndNames(t *testing.T) {
	rows := rowsOf(t, []string{"b", "a"}, []string{"1", "2"}, []string{"x", "y"})
	if len(rows) != 2 || rows[0][0] != "b" || rows[1][0] != "a" {
		t.Fatalf("names = %q, want b then a", rows)
	}
	for _, r := range rows {
		if len(r) != len(Header) {
			t.Errorf("row %q has %d fields, want %d", r, len(r), len(Header))
		}
	}
}

func TestRowsHeaderOnlyTable(t *testing.T) {
	rows := Rows(csvdata.Table{Header: []string{"x"}})
	want := [][]string{{"x", "empty", "0", "0", "0", "", "", "", "", ""}}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("Rows(header only) = %q, want %q", rows, want)
	}
}

func TestCountsRowsNullsDistinct(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		wantRows string
		wantNull string
		wantDist string
	}{
		{name: "one value", values: []string{"1"}, wantRows: "1", wantNull: "0", wantDist: "1"},
		{name: "repeats", values: []string{"7", "7", "7"}, wantRows: "3", wantNull: "0", wantDist: "1"},
		{name: "nulls excluded from distinct", values: []string{"x", "", "NA", "N/A", "NULL", "x"}, wantRows: "6", wantNull: "4", wantDist: "1"},
		{name: "case is distinct", values: []string{"Abc", "abc", "ABC"}, wantRows: "3", wantNull: "0", wantDist: "3"},
		{name: "numeric forms are distinct", values: []string{"1", "01", "1.0"}, wantRows: "3", wantNull: "0", wantDist: "3"},
		{name: "lookalikes are not null", values: []string{"0", "none", "nan"}, wantRows: "3", wantNull: "0", wantDist: "3"},
		{name: "all null", values: []string{"", "na", "null"}, wantRows: "3", wantNull: "3", wantDist: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := rowOf(t, tt.values...)
			if row[2] != tt.wantRows || row[3] != tt.wantNull || row[4] != tt.wantDist {
				t.Errorf("rows,nulls,distinct = %s,%s,%s; want %s,%s,%s",
					row[2], row[3], row[4], tt.wantRows, tt.wantNull, tt.wantDist)
			}
		})
	}
}

func TestIntegerExtremesAreNumeric(t *testing.T) {
	row := rowOf(t, "9", "10", "-3", "100", "NA")
	if row[1] != "integer" {
		t.Fatalf("type = %q, want integer", row[1])
	}
	if row[5] != "-3" || row[6] != "100" {
		t.Errorf("min,max = %s,%s; want -3,100 (numeric, not lexical)", row[5], row[6])
	}
}

func TestIntegerSumMeanMedian(t *testing.T) {
	tests := []struct {
		name                       string
		values                     []string
		wantSum, wantMean, wantMed string
	}{
		{name: "single", values: []string{"5"}, wantSum: "5", wantMean: "5.00", wantMed: "5.00"},
		{name: "odd count", values: []string{"3", "1", "2"}, wantSum: "6", wantMean: "2.00", wantMed: "2.00"},
		{name: "even count averages middles", values: []string{"4", "1", "3", "2"}, wantSum: "10", wantMean: "2.50", wantMed: "2.50"},
		{name: "numeric median order", values: []string{"10", "9", "100", "2"}, wantSum: "121", wantMean: "30.25", wantMed: "9.50"},
		{name: "negative mean rounds half away", values: []string{"-1", "0", "0", "0", "0", "0", "0", "0"}, wantSum: "-1", wantMean: "-0.13", wantMed: "0.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := rowOf(t, tt.values...)
			if row[7] != tt.wantSum || row[8] != tt.wantMean || row[9] != tt.wantMed {
				t.Errorf("sum,mean,median = %s,%s,%s; want %s,%s,%s",
					row[7], row[8], row[9], tt.wantSum, tt.wantMean, tt.wantMed)
			}
		})
	}
}

func TestIntegerSumIsExactBeyondFloat(t *testing.T) {
	row := rowOf(t, "9007199254740993", "1")
	if row[7] != "9007199254740994" {
		t.Errorf("sum = %s, want 9007199254740994", row[7])
	}
}

func TestNumberSumMeanMedianExact(t *testing.T) {
	tests := []struct {
		name                       string
		values                     []string
		wantSum, wantMean, wantMed string
	}{
		{name: "binary float trap", values: []string{"0.1", "0.2", "0.3", "0.7", "0.6"}, wantSum: "1.90", wantMean: "0.38", wantMed: "0.30"},
		{name: "even median tie rounds up", values: []string{"0.00", "0.01"}, wantSum: "0.01", wantMean: "0.01", wantMed: "0.01"},
		{name: "even median tie rounds away from zero", values: []string{"-0.01", "0.00"}, wantSum: "-0.01", wantMean: "-0.01", wantMed: "-0.01"},
		{name: "mixed integer and decimal", values: []string{"1", "2.5", "-3.25", "10.10", "4"}, wantSum: "14.35", wantMean: "2.87", wantMed: "2.50"},
		{name: "sum cancelling to zero", values: []string{"-0.001", "0.001"}, wantSum: "0.00", wantMean: "0.00", wantMed: "0.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := rowOf(t, tt.values...)
			if row[1] != "number" {
				t.Fatalf("type = %q, want number", row[1])
			}
			if row[7] != tt.wantSum || row[8] != tt.wantMean || row[9] != tt.wantMed {
				t.Errorf("sum,mean,median = %s,%s,%s; want %s,%s,%s",
					row[7], row[8], row[9], tt.wantSum, tt.wantMean, tt.wantMed)
			}
		})
	}
}

func TestNumberExtremesWrittenWithTwoDecimals(t *testing.T) {
	row := rowOf(t, "10.125", "-2.25", "0.0", "NULL")
	if row[5] != "-2.25" || row[6] != "10.13" {
		t.Errorf("min,max = %s,%s; want -2.25,10.13", row[5], row[6])
	}
}

func TestNumberNegativeZeroIsWrittenZero(t *testing.T) {
	row := rowOf(t, "-0.0")
	if row[5] != "0.00" || row[6] != "0.00" || row[7] != "0.00" {
		t.Errorf("min,max,sum = %s,%s,%s; want 0.00 each", row[5], row[6], row[7])
	}
}

func TestBigExactNumberSum(t *testing.T) {
	row := rowOf(t, "12345678901234567890.12", "0.01")
	if row[7] != "12345678901234567890.13" {
		t.Errorf("sum = %s, want 12345678901234567890.13", row[7])
	}
}

func TestDateExtremes(t *testing.T) {
	row := rowOf(t, "2025-03-01", "2024-12-31", "", "2025-01-15", "NA")
	if row[1] != "date" {
		t.Fatalf("type = %q, want date", row[1])
	}
	if row[5] != "2024-12-31" || row[6] != "2025-03-01" {
		t.Errorf("min,max = %s,%s; want 2024-12-31,2025-03-01", row[5], row[6])
	}
	if row[7] != "" || row[8] != "" || row[9] != "" {
		t.Errorf("sum,mean,median = %q,%q,%q; want empty", row[7], row[8], row[9])
	}
}

func TestBooleanExtremesLowerCase(t *testing.T) {
	tests := []struct {
		name    string
		values  []string
		wantMin string
		wantMax string
	}{
		{name: "all true", values: []string{"true", "TRUE", "True"}, wantMin: "true", wantMax: "true"},
		{name: "all false", values: []string{"FALSE", "false", ""}, wantMin: "false", wantMax: "false"},
		{name: "mixed", values: []string{"true", "False", "NA"}, wantMin: "false", wantMax: "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := rowOf(t, tt.values...)
			if row[1] != "boolean" {
				t.Fatalf("type = %q, want boolean", row[1])
			}
			if row[5] != tt.wantMin || row[6] != tt.wantMax {
				t.Errorf("min,max = %s,%s; want %s,%s", row[5], row[6], tt.wantMin, tt.wantMax)
			}
		})
	}
}

func TestTextExtremesByByteValue(t *testing.T) {
	row := rowOf(t, "b", "B", "a", "Z", "NA", "pad")
	if row[1] != "text" {
		t.Fatalf("type = %q, want text", row[1])
	}
	if row[5] != "B" || row[6] != "pad" {
		t.Errorf("min,max = %s,%s; want B,pad", row[5], row[6])
	}
	row = rowOf(t, "10", "9", "x")
	if row[1] != "text" || row[5] != "10" || row[6] != "x" {
		t.Errorf("row = %q; want text with min 10 and max x", row)
	}
}

func TestEmptyColumnHasNoStatistics(t *testing.T) {
	row := rowOf(t, "", "NA", "NULL")
	want := []string{"c", "empty", "3", "3", "0", "", "", "", "", ""}
	if !reflect.DeepEqual(row, want) {
		t.Errorf("row = %q, want %q", row, want)
	}
}

func TestTextAndDateLeaveNumericFieldsEmpty(t *testing.T) {
	for _, values := range [][]string{{"abc", "1"}, {"2025-01-01", "2025-01-02"}, {"true", "false"}} {
		row := rowOf(t, values...)
		if row[7] != "" || row[8] != "" || row[9] != "" {
			t.Errorf("values %q: sum,mean,median = %q,%q,%q; want empty", values, row[7], row[8], row[9])
		}
	}
}

func TestNumericFieldsSkipNulls(t *testing.T) {
	row := rowOf(t, "2", "", "NULL", "4")
	if row[7] != "6" || row[8] != "3.00" || row[9] != "3.00" {
		t.Errorf("sum,mean,median = %s,%s,%s; want 6,3.00,3.00", row[7], row[8], row[9])
	}
}
