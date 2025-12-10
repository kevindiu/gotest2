package utils

import (
	"strings"
	"unicode"
)

// ParseISBN checks if a string is a valid ISBN-10-ish format for our example.
// It basically checks if it has 10 digits or 9 digits and an 'X'.
// This is a simple implementation for fuzz testing demonstration.
func ParseISBN(isbn string) bool {
	// Remove hyphens
	s := strings.ReplaceAll(isbn, "-", "")
	if len(s) != 10 {
		return false
	}

	for i, r := range s {
		if i == 9 {
			if !unicode.IsDigit(r) && r != 'X' && r != 'x' {
				return false
			}
		} else {
			if !unicode.IsDigit(r) {
				return false
			}
		}
	}
	return true
}

// FormatISBN adds hyphens to a raw 10-char ISBN string.
// Example: 123456789X -> 1-234-56789-X
// This returns the original string if it's not length 10.
func FormatISBN(raw string) string {
	if len(raw) != 10 {
		return raw
	}
	return raw[0:1] + "-" + raw[1:4] + "-" + raw[4:9] + "-" + raw[9:]
}
