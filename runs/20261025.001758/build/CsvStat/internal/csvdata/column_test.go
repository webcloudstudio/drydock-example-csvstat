package csvdata

import "testing"

func TestColumnFindsExactName(t *testing.T) {
	table := Table{Header: []string{"id", "qty", "Qty"}}
	tests := []struct {
		name      string
		wantIndex int
		wantFound bool
	}{
		{name: "id", wantIndex: 0, wantFound: true},
		{name: "qty", wantIndex: 1, wantFound: true},
		{name: "Qty", wantIndex: 2, wantFound: true},
		{name: "QTY", wantFound: false},
		{name: "nosuch", wantFound: false},
		{name: "", wantFound: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, found := Column(table, tt.name)
			if found != tt.wantFound {
				t.Fatalf("Column(%q) found = %v, want %v", tt.name, found, tt.wantFound)
			}
			if found && index != tt.wantIndex {
				t.Errorf("Column(%q) index = %d, want %d", tt.name, index, tt.wantIndex)
			}
		})
	}
}

func TestColumnOnHeaderOnlyTable(t *testing.T) {
	if _, found := Column(Table{}, "id"); found {
		t.Error("Column(empty table) found a name, want none")
	}
}
