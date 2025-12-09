package main

import (
	"testing"
)

func FuzzMultiError(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, n int) {
		MultiError(
			n,
		)
	})
}
