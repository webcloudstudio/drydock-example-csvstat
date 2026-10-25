package csvdata

// Column looks up a header name by exact string equality with the trimmed header
// names. It returns the column index and whether the name was found.
func Column(t Table, name string) (int, bool) {
	for i, header := range t.Header {
		if header == name {
			return i, true
		}
	}
	return 0, false
}
