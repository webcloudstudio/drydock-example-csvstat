package csvdata

import "testing"

func TestIsNullEmptyAndSentinels(t *testing.T) {
	values := []string{"", "NA", "na", "Na", "N/A", "n/a", "N/a", "NULL", "null", "Null"}
	for _, v := range values {
		if !IsNull(v) {
			t.Errorf("IsNull(%q) = false, want true", v)
		}
	}
}

func TestIsNullRejectsLookalikes(t *testing.T) {
	values := []string{"0", "nan", "none", "NAN", "nul", "n/a/", "N A", "NA1", "-"}
	for _, v := range values {
		if IsNull(v) {
			t.Errorf("IsNull(%q) = true, want false", v)
		}
	}
}
