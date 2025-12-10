package main

import (
	"testing"
)

func Benchmark_main(b *testing.B) {
	// No parameters, simple benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		main()
	}
}

func Benchmark_getPort(b *testing.B) {
	// No parameters, simple benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getPort()
	}
}
