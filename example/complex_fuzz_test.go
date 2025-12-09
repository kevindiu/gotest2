package main

import (
	"testing"
)

func FuzzParseConfig(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, data []byte) {
		ParseConfig(
			data,
		)
	})
}
