package main

import (
	"testing"
)

func FuzzReverse(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, s string) {
		Reverse(
			s,
		)
	})
}
