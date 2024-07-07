package list

import (
	"strings"
)

// FromStringP - Given a string pointer, parse by delimiter and clean up
func FromStringP(s *string, sep string, pruneQuotes bool) (data []string) {

	// Set up our cleaner function, so we are only doing our if pruneQuotes check one time.
	var cleaner func(*string) *string

	if s == nil || *s == "" {
		return data
	}

	if pruneQuotes {
		cleaner = func(i *string) *string {
			*i = strings.TrimSpace(*i)
			*i = strings.TrimPrefix(*i, singleQuote)
			*i = strings.TrimPrefix(*i, doubleQuote)
			*i = strings.TrimSuffix(*i, singleQuote)
			*i = strings.TrimSuffix(*i, doubleQuote)
			return i
		}
	} else {
		cleaner = func(i *string) *string {
			return i
		}
	}

	data = strings.Split(*(cleaner(s)), sep)

	for i := range data {
		// Creating a pointer because we want to defeat multiple copies of the string
		dp := &data[i]
		dp = cleaner(dp)
	}

	return data

}
