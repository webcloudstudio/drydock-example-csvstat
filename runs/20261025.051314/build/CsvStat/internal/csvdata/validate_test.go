package csvdata

import (
	"errors"
	"strings"
	"testing"
)

// checkInput runs the same pipeline the cli uses: Read, Trim, Validate.
func checkInput(t *testing.T, input string) error {
	t.Helper()
	table, err := Read(strings.NewReader(input))
	if err != nil {
		return err
	}
	Trim(table)
	return Validate(table)
}

func TestRejectEmptyFile(t *testing.T) {
	if err := checkInput(t, ""); !errors.Is(err, ErrRejected) {
		t.Fatalf("empty file error = %v, want ErrRejected", err)
	}
}

func TestRejectBlankLinesOnly(t *testing.T) {
	if err := checkInput(t, "\n\n\n"); !errors.Is(err, ErrRejected) {
		t.Fatalf("blank-lines-only error = %v, want ErrRejected", err)
	}
}

func TestRejectEmptyHeaderName(t *testing.T) {
	for _, input := range []string{"id,,name\n1,x,y\n", "id,  ,name\n1,x,y\n"} {
		if err := checkInput(t, input); !errors.Is(err, ErrRejected) {
			t.Errorf("checkInput(%q) error = %v, want ErrRejected", input, err)
		}
	}
}

func TestRejectRepeatedHeaderName(t *testing.T) {
	for _, input := range []string{"id,name,id\n1,x,2\n", "id, name ,name\n1,x,2\n"} {
		if err := checkInput(t, input); !errors.Is(err, ErrRejected) {
			t.Errorf("checkInput(%q) error = %v, want ErrRejected", input, err)
		}
	}
}

func TestRepeatedHeaderNameIsCaseSensitive(t *testing.T) {
	if err := checkInput(t, "id,name,ID\n1,x,2\n"); err != nil {
		t.Fatalf("distinct-case headers rejected: %v", err)
	}
}

func TestRejectRaggedRecord(t *testing.T) {
	err := checkInput(t, "a,b\n1,2\n3,4,5\n6,7\n")

	if !errors.Is(err, ErrRejected) {
		t.Fatalf("ragged record error = %v, want ErrRejected", err)
	}
	var rejected *RejectedError
	if !errors.As(err, &rejected) {
		t.Fatalf("ragged record error = %v, want *RejectedError", err)
	}
	if rejected.Line != 3 {
		t.Fatalf("RejectedError.Line = %d, want 3", rejected.Line)
	}
}

func TestHeaderOnlyFileAccepted(t *testing.T) {
	if err := checkInput(t, "a,b\n"); err != nil {
		t.Fatalf("header-only file rejected: %v", err)
	}
}

func TestTrailingNewlineAddsNoRecord(t *testing.T) {
	for _, input := range []string{"a,b\n1,2\n", "a,b\n1,2"} {
		if err := checkInput(t, input); err != nil {
			t.Errorf("checkInput(%q) rejected: %v", input, err)
		}
		table := mustRead(t, input)
		if len(table.Records) != 1 {
			t.Errorf("Read(%q) got %d records, want 1", input, len(table.Records))
		}
	}
}
