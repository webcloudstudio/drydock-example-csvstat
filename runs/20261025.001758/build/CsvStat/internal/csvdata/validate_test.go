package csvdata

import (
	"errors"
	"testing"
)

func TestValidateAcceptsHeaderOnly(t *testing.T) {
	table := Table{Header: []string{"sku", "price"}}
	if err := Validate(table); err != nil {
		t.Errorf("Validate(header only) error = %v, want nil", err)
	}
}

func TestValidateAcceptsWellFormedTable(t *testing.T) {
	table := Table{
		Header:  []string{"a", "b"},
		Records: []Record{{Line: 2, Fields: []string{"1", "2"}}},
	}
	if err := Validate(table); err != nil {
		t.Errorf("Validate(well formed) error = %v, want nil", err)
	}
}

func TestValidateRejectsWithRejectError(t *testing.T) {
	tests := []struct {
		name     string
		table    Table
		wantLine int
	}{
		{name: "no records", table: Table{}, wantLine: 0},
		{name: "empty header name", table: Table{Header: []string{"a", ""}}, wantLine: 0},
		{name: "repeated header name", table: Table{Header: []string{"a", "b", "a"}}, wantLine: 0},
		{
			name:     "short record",
			table:    Table{Header: []string{"a", "b", "c"}, Records: []Record{{Line: 2, Fields: []string{"1", "2"}}}},
			wantLine: 2,
		},
		{
			name:     "long record after valid record",
			table:    Table{Header: []string{"a", "b"}, Records: []Record{{Line: 2, Fields: []string{"1", "2"}}, {Line: 5, Fields: []string{"1", "2", "3"}}}},
			wantLine: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.table)

			var reject *RejectError
			if !errors.As(err, &reject) {
				t.Fatalf("Validate() error = %v, want *RejectError", err)
			}
			if reject.Line != tt.wantLine {
				t.Errorf("RejectError.Line = %d, want %d", reject.Line, tt.wantLine)
			}
		})
	}
}

func TestValidateRepeatedHeadersAreCaseSensitive(t *testing.T) {
	table := Table{Header: []string{"a", "A"}}
	if err := Validate(table); err != nil {
		t.Errorf("Validate(%q) error = %v, want nil; names differing in case are distinct", table.Header, err)
	}
}
