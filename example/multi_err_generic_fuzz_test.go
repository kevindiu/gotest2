package main

import (
	"testing"
)

func FuzzGenericMultiError(f *testing.F) {
	// TODO: Initialize concrete type for T
	type T = int
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, val T, n int) {
		GenericMultiError(
			val,
			n,
		)
	})
}
