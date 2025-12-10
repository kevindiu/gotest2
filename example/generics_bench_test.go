package main

import (
	"testing"
)

func BenchmarkGenericSum(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkGenericSum[int](b, []testCaseBenchmarkGenericSum[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkMyList_Add(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMyList_Add[int](b, []testCaseBenchmarkMyList_Add[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkMyList_Get(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMyList_Get[int](b, []testCaseBenchmarkMyList_Get[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkSwap(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkSwap[int](b, []testCaseBenchmarkSwap[int]{
			// TODO: Add benchmark cases
		})
	})
}

func benchmarkGenericSum[T int | float64](b *testing.B, tests []testCaseBenchmarkGenericSum[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkGenericSum[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkGenericSum[T]) {}
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
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				GenericSum[T](
					bb.args.a,
					bb.args.b,
				)
			}
		})
	}
}

func benchmarkMyList_Add[T any](b *testing.B, tests []testCaseBenchmarkMyList_Add[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMyList_Add[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMyList_Add[T]) {}
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
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bb.receiver.Add(
					bb.args.item,
				)
			}
		})
	}
}

func benchmarkMyList_Get[T any](b *testing.B, tests []testCaseBenchmarkMyList_Get[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMyList_Get[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMyList_Get[T]) {}
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
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bb.receiver.Get(
					bb.args.index,
				)
			}
		})
	}
}

func benchmarkSwap[T any](b *testing.B, tests []testCaseBenchmarkSwap[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkSwap[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkSwap[T]) {}
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
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Swap[T](
					bb.args.a,
					bb.args.b,
				)
			}
		})
	}
}

type testCaseBenchmarkGenericSum[T int | float64] struct {
	name string
	args struct {
		a T
		b T
	}
	init    func(b *testing.B, tt *testCaseBenchmarkGenericSum[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkGenericSum[T])
}

type testCaseBenchmarkMyList_Add[T any] struct {
	name string
	args struct {
		item T
	}
	receiver *MyList[T]
	init     func(b *testing.B, tt *testCaseBenchmarkMyList_Add[T])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMyList_Add[T])
}

type testCaseBenchmarkMyList_Get[T any] struct {
	name string
	args struct {
		index int
	}
	receiver *MyList[T]
	init     func(b *testing.B, tt *testCaseBenchmarkMyList_Get[T])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMyList_Get[T])
}

type testCaseBenchmarkSwap[T any] struct {
	name string
	args struct {
		a T
		b T
	}
	init    func(b *testing.B, tt *testCaseBenchmarkSwap[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkSwap[T])
}
