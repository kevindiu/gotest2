package utils

import (
	"testing"
)

func FuzzParseISBN(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, isbn string) {
		ParseISBN(
			isbn,
		)
	})
}

func FuzzFormatISBN(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, raw string) {
		FormatISBN(
			raw,
		)
	})
}
