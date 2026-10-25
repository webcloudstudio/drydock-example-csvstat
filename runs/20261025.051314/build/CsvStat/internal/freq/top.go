package freq

// Top returns the first n entries of an ordered listing. A limit of n or more returns every entry,
// and a limit below zero returns none. Call it after Count, so the limit applies after ordering.
func Top(entries []Entry, n int) []Entry {
	if n < 0 {
		n = 0
	}
	if n >= len(entries) {
		return entries
	}
	return entries[:n]
}
