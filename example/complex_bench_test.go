package main

import (
	"testing"
)

func BenchmarkParseConfig(b *testing.B) {
	type args struct {
		data []byte
	}
	type test struct {
		name    string
		args    args
		init    func(b *testing.B, tt *test)
		cleanup func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			if bb.init == nil {
				bb.init = defaultInit
			}
			bb.init(b, &bb)
			if bb.cleanup == nil {
				bb.cleanup = defaultCleanup
			}
			defer bb.cleanup(b, &bb)
			for i := 0; i < b.N; i++ {
				ParseConfig(
					bb.args.data,
				)
			}
		})
	}
}
