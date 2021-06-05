package main

import (
	"strconv"
	"strings"
)

// Parse a base 10 integer from the provided string, trimming any whitespace
// Returns the parsed integer and ok signal
func parseInt64(s string) (int64, bool) {
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}

// Parse a floating point number from the provided string, trimming any whitespace
func parseFloat64(s string) (float64, bool) {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, false
	}
	return f, true
}
