package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenericSum(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestGenericSum[int](t, []testGenericSumTestCase[int]{
			{
				name: "Positive int",
				args: struct {
					a int
					b int
				}{a: 1, b: 2},
				want: testGenericSumWants[int]{want0: 3},
			},
			{
				name: "Negative int",
				args: struct {
					a int
					b int
				}{a: -5, b: -5},
				want: testGenericSumWants[int]{want0: -10},
			},
		})
	})
	t.Run("float64", func(t *testing.T) {
		runTestGenericSum[float64](t, []testGenericSumTestCase[float64]{
			{
				name: "Positive float",
				args: struct {
					a float64
					b float64
				}{a: 1.5, b: 2.5},
				want: testGenericSumWants[float64]{want0: 4.0},
			},
			{
				name: "Mixed float",
				args: struct {
					a float64
					b float64
				}{a: -1.5, b: 2.5},
				want: testGenericSumWants[float64]{want0: 1.0},
			},
		})
	})
}

func TestMyList_Add(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestMyList_Add[int](t, []testMyList_AddTestCase[int]{
			{
				name:     "Add to empty list",
				receiver: &MyList[int]{items: []int{}},
				args: struct {
					item int
				}{item: 10},
				Validate: func(t *testing.T, tt *testMyList_AddTestCase[int]) error {
					if len(tt.receiver.items) != 1 {
						return fmt.Errorf("expected length 1, got %d", len(tt.receiver.items))
					}
					if tt.receiver.items[0] != 10 {
						return fmt.Errorf("expected item 10, got %d", tt.receiver.items[0])
					}
					return nil
				},
			},
			{
				name:     "Append to existing",
				receiver: &MyList[int]{items: []int{1, 2}},
				args: struct {
					item int
				}{item: 3},
				Validate: func(t *testing.T, tt *testMyList_AddTestCase[int]) error {
					if len(tt.receiver.items) != 3 {
						return fmt.Errorf("expected length 3, got %d", len(tt.receiver.items))
					}
					return nil
				},
			},
		})
	})
	t.Run("string", func(t *testing.T) {
		runTestMyList_Add[string](t, []testMyList_AddTestCase[string]{
			{
				name:     "Add string",
				receiver: &MyList[string]{items: []string{"hello"}},
				args: struct {
					item string
				}{item: "world"},
				Validate: func(t *testing.T, tt *testMyList_AddTestCase[string]) error {
					if len(tt.receiver.items) != 2 {
						return fmt.Errorf("expected length 2, got %d", len(tt.receiver.items))
					}
					if tt.receiver.items[1] != "world" {
						return fmt.Errorf("expected item 'world', got %s", tt.receiver.items[1])
					}
					return nil
				},
			},
		})
	})
}

func TestMyList_Get(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestMyList_Get[int](t, []testMyList_GetTestCase[int]{
			{
				name:     "Get first item",
				receiver: &MyList[int]{items: []int{10, 20, 30}},
				args: struct {
					index int
				}{index: 0},
				want: testMyList_GetWants[int]{want0: 10},
			},
			{
				name:     "Get last item",
				receiver: &MyList[int]{items: []int{10, 20, 30}},
				args: struct {
					index int
				}{index: 2},
				want: testMyList_GetWants[int]{want0: 30},
			},
		})
	})
}

func TestSwap(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestSwap[int](t, []testSwapTestCase[int]{
			{
				name: "Swap integers",
				args: struct {
					a int
					b int
				}{a: 1, b: 2},
				want: testSwapWants[int]{want0: 2, want1: 1},
			},
			{
				name: "Swap same values",
				args: struct {
					a int
					b int
				}{a: 5, b: 5},
				want: testSwapWants[int]{want0: 5, want1: 5},
			},
		})
	})
	t.Run("string", func(t *testing.T) {
		runTestSwap[string](t, []testSwapTestCase[string]{
			{
				name: "Swap strings",
				args: struct {
					a string
					b string
				}{a: "foo", b: "bar"},
				want: testSwapWants[string]{want0: "bar", want1: "foo"},
			},
		})
	})
}

func runTestGenericSum[T int | float64](t *testing.T, cases []testGenericSumTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testGenericSumTestCase[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericSum() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0 := GenericSum(
				tt.args.a,
				tt.args.b,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("GenericSum() validation failed: %v", err)
			}
		})
	}
}

func runTestMyList_Add[T any](t *testing.T, cases []testMyList_AddTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, tt *testMyList_AddTestCase[T]) error {
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			tt.receiver.Add(
				tt.args.item,
			)

			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, &tt); err != nil {
				t.Errorf("MyList_Add() validation failed: %v", err)
			}
		})
	}
}

func runTestMyList_Get[T any](t *testing.T, cases []testMyList_GetTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testMyList_GetTestCase[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("MyList_Get() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0 := tt.receiver.Get(
				tt.args.index,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("MyList_Get() validation failed: %v", err)
			}
		})
	}
}

func runTestSwap[T any](t *testing.T, cases []testSwapTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, got1 T, tt *testSwapTestCase[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("Swap() got0 = %v, want %v", got0, tt.want.want0)
			}
			if !reflect.DeepEqual(got1, tt.want.want1) {
				return fmt.Errorf("Swap() got1 = %v, want %v", got1, tt.want.want1)
			}
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0, got1 := Swap(
				tt.args.a,
				tt.args.b,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, got1, &tt); err != nil {
				t.Errorf("Swap() validation failed: %v", err)
			}
		})
	}
}

type testGenericSumWants[T int | float64] struct {
	want0 T
}

type testGenericSumTestCase[T int | float64] struct {
	name string
	args struct {
		a T
		b T
	}
	want testGenericSumWants[T]

	Init     func(t *testing.T, tt *testGenericSumTestCase[T])
	Cleanup  func(t *testing.T, tt *testGenericSumTestCase[T])
	Validate func(t *testing.T, got0 T, tt *testGenericSumTestCase[T]) error
}

type testMyList_AddTestCase[T any] struct {
	name     string
	receiver *MyList[T]
	args     struct {
		item T
	}

	Init     func(t *testing.T, tt *testMyList_AddTestCase[T])
	Cleanup  func(t *testing.T, tt *testMyList_AddTestCase[T])
	Validate func(t *testing.T, tt *testMyList_AddTestCase[T]) error
}

type testMyList_GetWants[T any] struct {
	want0 T
}

type testMyList_GetTestCase[T any] struct {
	name     string
	receiver *MyList[T]
	args     struct {
		index int
	}
	want testMyList_GetWants[T]

	Init     func(t *testing.T, tt *testMyList_GetTestCase[T])
	Cleanup  func(t *testing.T, tt *testMyList_GetTestCase[T])
	Validate func(t *testing.T, got0 T, tt *testMyList_GetTestCase[T]) error
}

type testSwapWants[T any] struct {
	want0 T
	want1 T
}

type testSwapTestCase[T any] struct {
	name string
	args struct {
		a T
		b T
	}
	want testSwapWants[T]

	Init     func(t *testing.T, tt *testSwapTestCase[T])
	Cleanup  func(t *testing.T, tt *testSwapTestCase[T])
	Validate func(t *testing.T, got0 T, got1 T, tt *testSwapTestCase[T]) error
}
