package main

import (
	"testing"
)

func FuzzGenericAdd(f *testing.F) {
	// TODO: Initialize concrete type for T
	type T = int
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a T, b T) {
		GenericAdd(
			a,
			b,
		)
	})
}
