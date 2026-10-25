package freq

import (
	"reflect"
	"testing"
)

var threeEntries = []Entry{{"c", 5}, {"b", 3}, {"a", 1}}

func TestTopTakesFirstN(t *testing.T) {
	got := Top(threeEntries, 2)
	if want := threeEntries[:2]; !reflect.DeepEqual(got, want) {
		t.Fatalf("Top(2) = %+v, want %+v", got, want)
	}
}

func TestTopLargerThanListingListsAll(t *testing.T) {
	if got := Top(threeEntries, 50); !reflect.DeepEqual(got, threeEntries) {
		t.Fatalf("Top(50) = %+v, want all entries", got)
	}
}

func TestTopZeroListsNone(t *testing.T) {
	if got := Top(threeEntries, 0); len(got) != 0 {
		t.Fatalf("Top(0) = %+v, want none", got)
	}
}

func TestTopNegativeListsNone(t *testing.T) {
	if got := Top(threeEntries, -1); len(got) != 0 {
		t.Fatalf("Top(-1) = %+v, want none", got)
	}
}

func TestTopOfEmptyListing(t *testing.T) {
	if got := Top(nil, 5); len(got) != 0 {
		t.Fatalf("Top(nil, 5) = %+v, want none", got)
	}
}
