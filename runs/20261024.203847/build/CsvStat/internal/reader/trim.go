package reader

import "strings"

// Trim removes leading and trailing space characters (U+0020) only.
func Trim(s string) string {
	return strings.Trim(s, " ")
}

// IsNull reports whether an already trimmed value is null: empty, or NA, N/A, or NULL
// ignoring case. Every other value is non-null.
func IsNull(s string) bool {
	return s == "" ||
		strings.EqualFold(s, "NA") ||
		strings.EqualFold(s, "N/A") ||
		strings.EqualFold(s, "NULL")
}

// TrimTable returns a copy of t with every header name and every value trimmed. Lines is kept.
func TrimTable(t Table) Table {
	out := Table{
		Header: trimAll(t.Header),
		Rows:   make([][]string, len(t.Rows)),
		Lines:  append([]int(nil), t.Lines...),
	}
	for i, row := range t.Rows {
		out.Rows[i] = trimAll(row)
	}
	return out
}

func trimAll(values []string) []string {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = Trim(v)
	}
	return out
}
