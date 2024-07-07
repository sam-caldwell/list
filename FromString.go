package list

// FromString - Given a string, parse by delimiter and clean up
func FromString(s string, sep string, pruneQuotes bool) []string {
	return FromStringP(&s, sep, pruneQuotes)
}
