package csvdata

import "testing"

func TestIsNullEmpty(t *testing.T) {
	if !IsNull("") {
		t.Fatal(`IsNull("") = false, want true`)
	}
}

func TestIsNullSpellings(t *testing.T) {
	for _, value := range []string{"NA", "na", "Na", "N/A", "n/a", "NULL", "null", "Null"} {
		if !IsNull(value) {
			t.Errorf("IsNull(%q) = false, want true", value)
		}
	}
}

func TestIsNullNonNullValues(t *testing.T) {
	for _, value := range []string{"0", "false", "NAN", "N/B", "NULLS", "NA1", "N A", "nil", "none", "-"} {
		if IsNull(value) {
			t.Errorf("IsNull(%q) = true, want false", value)
		}
	}
}

func TestQuotedEmptyFieldIsNull(t *testing.T) {
	table := mustRead(t, "a,b\n\"\",x\n")

	if !IsNull(table.Records[0].Fields[0]) {
		t.Fatalf("quoted empty field %q is not null", table.Records[0].Fields[0])
	}
}

func TestSpacesOnlyFieldIsNull(t *testing.T) {
	table := mustRead(t, "a,b\n   ,x\n")
	Trim(table)

	if !IsNull(table.Records[0].Fields[0]) {
		t.Fatalf("spaces-only field %q is not null after trimming", table.Records[0].Fields[0])
	}
}
