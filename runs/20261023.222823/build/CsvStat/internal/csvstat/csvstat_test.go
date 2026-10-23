package csvstat

import (
	"strings"
	"testing"
)

func TestLoadSeparatesHeaderAndRecords(t *testing.T) {
	input, err := Load(strings.NewReader("name,age\nAda,36\n"))
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if len(input.Header) != 2 || input.Header[0] != "name" || input.Header[1] != "age" {
		t.Fatalf("header = %#v, want [name age]", input.Header)
	}
	if len(input.Records) != 1 || input.Records[0][0] != "Ada" {
		t.Fatalf("records = %#v, want one Ada record", input.Records)
	}
}

func TestLoadRejectsEmptyInput(t *testing.T) {
	_, err := Load(strings.NewReader(""))
	if err == nil {
		t.Fatal("Load(empty) returned nil error")
	}
}
