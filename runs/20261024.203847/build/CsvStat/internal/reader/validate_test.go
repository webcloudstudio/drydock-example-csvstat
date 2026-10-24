package reader

import (
	"errors"
	"strings"
	"testing"
)

func validated(t *testing.T, in string) error {
	t.Helper()
	tbl, err := Read(strings.NewReader(in))
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	return Validate(TrimTable(tbl))
}

func rejectOf(t *testing.T, err error) *RejectError {
	t.Helper()
	var rej *RejectError
	if !errors.As(err, &rej) {
		t.Fatalf("Validate error = %v, want *RejectError", err)
	}
	return rej
}

func TestValidateEmptyFile(t *testing.T) {
	rejectOf(t, validated(t, ""))
}

func TestValidateEmptyHeaderName(t *testing.T) {
	rejectOf(t, validated(t, "a,,c\n1,2,3\n"))
	rejectOf(t, validated(t, "a, ,c\n1,2,3\n"))
}

func TestValidateRepeatedHeaderName(t *testing.T) {
	rejectOf(t, validated(t, "a,b,a\n1,2,3\n"))
	rejectOf(t, validated(t, "a,b, a \n1,2,3\n"))
}

func TestValidateRaggedRecordLine(t *testing.T) {
	rej := rejectOf(t, validated(t, "a,b\n1,2\n3\n"))
	if rej.Line != 3 {
		t.Errorf("Line = %d, want 3", rej.Line)
	}
}

func TestValidateShortAndLongRecords(t *testing.T) {
	rejectOf(t, validated(t, "a,b,c\n1,2,3\n4,5\n"))
	rejectOf(t, validated(t, "a,b,c\n1,2,3\n4,5,6,7\n"))
}

func TestValidateHeaderOnlyAccepted(t *testing.T) {
	if err := validated(t, "a,b\n"); err != nil {
		t.Errorf("Validate(header only) = %v, want nil", err)
	}
}

func TestValidateAcceptsWellFormed(t *testing.T) {
	if err := validated(t, "a,b\n1,2\n"); err != nil {
		t.Errorf("Validate = %v, want nil", err)
	}
}
