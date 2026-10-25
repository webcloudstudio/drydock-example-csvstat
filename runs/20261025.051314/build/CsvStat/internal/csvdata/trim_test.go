package csvdata

import (
	"reflect"
	"testing"
)

func TestTrimHeaderNames(t *testing.T) {
	table := mustRead(t, "  name ,  qty\n1,2\n")
	Trim(table)

	if want := []string{"name", "qty"}; !reflect.DeepEqual(table.Header, want) {
		t.Fatalf("Header = %q, want %q", table.Header, want)
	}
}

func TestTrimValues(t *testing.T) {
	table := mustRead(t, "a,b\n  Globex  , Acme Corp \n")
	Trim(table)

	if want := []string{"Globex", "Acme Corp"}; !reflect.DeepEqual(table.Records[0].Fields, want) {
		t.Fatalf("Fields = %q, want %q", table.Records[0].Fields, want)
	}
}

func TestTrimQuotedPadding(t *testing.T) {
	table := mustRead(t, "a\n\"  Acme, Inc.  \"\n")
	Trim(table)

	if want := []string{"Acme, Inc."}; !reflect.DeepEqual(table.Records[0].Fields, want) {
		t.Fatalf("Fields = %q, want %q", table.Records[0].Fields, want)
	}
}

func TestTrimOnlySpaces(t *testing.T) {
	// Only U+0020 is trimmed: the leading and trailing tabs stay, and the
	// interior space stays.
	table := mustRead(t, "  \ta b\t  \n \tv w\t \n")
	Trim(table)

	if want := []string{"\ta b\t"}; !reflect.DeepEqual(table.Header, want) {
		t.Fatalf("Header = %q, want %q", table.Header, want)
	}
	if want := []string{"\tv w\t"}; !reflect.DeepEqual(table.Records[0].Fields, want) {
		t.Fatalf("Fields = %q, want %q", table.Records[0].Fields, want)
	}
}

func TestTrimKeepsLines(t *testing.T) {
	table := mustRead(t, "a,b\n x , y \n\"p\nq\", z\n")
	beforeLines := []int{table.Records[0].Line, table.Records[1].Line}
	beforeCounts := []int{len(table.Records[0].Fields), len(table.Records[1].Fields)}

	Trim(table)

	afterLines := []int{table.Records[0].Line, table.Records[1].Line}
	afterCounts := []int{len(table.Records[0].Fields), len(table.Records[1].Fields)}
	if !reflect.DeepEqual(afterLines, beforeLines) {
		t.Fatalf("record lines = %v after Trim, want %v", afterLines, beforeLines)
	}
	if !reflect.DeepEqual(afterCounts, beforeCounts) {
		t.Fatalf("field counts = %v after Trim, want %v", afterCounts, beforeCounts)
	}
}
