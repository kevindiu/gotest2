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
		testGenericSum[int](t, []testCaseGenericSum[int]{
			// TODO: Add test cases.
		})
	})
}

func TestMyList_Add(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMyList_Add[int](t, []testCaseMyList_Add[int]{
			// TODO: Add test cases.
		})
	})
}

func TestMyList_Get(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMyList_Get[int](t, []testCaseMyList_Get[int]{
			// TODO: Add test cases.
		})
	})
}

func TestSwap(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testSwap[int](t, []testCaseSwap[int]{
			// TODO: Add test cases.
		})
	})
}

func testGenericSum[T int | float64](t *testing.T, cases []testCaseGenericSum[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testCaseGenericSum[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericSum() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseGenericSum[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseGenericSum[T]) {}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0 := GenericSum(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("GenericSum() validation failed: %v", err)
			}
		})
	}
}

func testMyList_Add[T any](t *testing.T, cases []testCaseMyList_Add[T]) {
	t.Parallel()
	for _, tt := range cases {

		defaultValidate := func(t *testing.T, tt *testCaseMyList_Add[T]) error {
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMyList_Add[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMyList_Add[T]) {}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			tt.receiver.Add(
				tt.args.item,
			)

			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("MyList_Add() validation failed: %v", err)
			}
		})
	}
}

func testMyList_Get[T any](t *testing.T, cases []testCaseMyList_Get[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testCaseMyList_Get[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("MyList_Get() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMyList_Get[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMyList_Get[T]) {}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0 := tt.receiver.Get(
				tt.args.index,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("MyList_Get() validation failed: %v", err)
			}
		})
	}
}

func testSwap[T any](t *testing.T, cases []testCaseSwap[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, got1 T, tt *testCaseSwap[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("Swap() got0 = %v, want %v", got0, tt.want.want0)
			}
			if !reflect.DeepEqual(got1, tt.want.want1) {
				return fmt.Errorf("Swap() got1 = %v, want %v", got1, tt.want.want1)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseSwap[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseSwap[T]) {}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0, got1 := Swap(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, got1, &tt); err != nil {
				t.Errorf("Swap() validation failed: %v", err)
			}
		})
	}
}

type testGenericSumWants[T int | float64] struct {
	want0 T
}

type testCaseGenericSum[T int | float64] struct {
	name string
	args struct {
		a T
		b T
	}
	want     testGenericSumWants[T]
	init     func(t *testing.T, tt *testCaseGenericSum[T])
	cleanup  func(t *testing.T, tt *testCaseGenericSum[T])
	validate func(t *testing.T, got0 T, tt *testCaseGenericSum[T]) error
}

type testCaseMyList_Add[T any] struct {
	name     string
	receiver *MyList[T]
	args     struct {
		item T
	}
	init     func(t *testing.T, tt *testCaseMyList_Add[T])
	cleanup  func(t *testing.T, tt *testCaseMyList_Add[T])
	validate func(t *testing.T, tt *testCaseMyList_Add[T]) error
}

type testMyList_GetWants[T any] struct {
	want0 T
}

type testCaseMyList_Get[T any] struct {
	name     string
	receiver *MyList[T]
	args     struct {
		index int
	}
	want     testMyList_GetWants[T]
	init     func(t *testing.T, tt *testCaseMyList_Get[T])
	cleanup  func(t *testing.T, tt *testCaseMyList_Get[T])
	validate func(t *testing.T, got0 T, tt *testCaseMyList_Get[T]) error
}

type testSwapWants[T any] struct {
	want0 T
	want1 T
}

type testCaseSwap[T any] struct {
	name string
	args struct {
		a T
		b T
	}
	want     testSwapWants[T]
	init     func(t *testing.T, tt *testCaseSwap[T])
	cleanup  func(t *testing.T, tt *testCaseSwap[T])
	validate func(t *testing.T, got0 T, got1 T, tt *testCaseSwap[T]) error
}
