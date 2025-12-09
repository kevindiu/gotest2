package main

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestGenericAdd(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testGenericAdd[int](t, []testCaseGenericAdd[int]{
			// TODO: Add test cases.
		})
	})
}

func TestGenericAppend(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testGenericAppend[int](t, []testCaseGenericAppend[int]{
			// TODO: Add test cases.
		})
	})
}

func testGenericAdd[T constraints.Ordered](t *testing.T, cases []testCaseGenericAdd[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testCaseGenericAdd[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericAdd() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseGenericAdd[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseGenericAdd[T]) {}
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
			got0 := GenericAdd(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("GenericAdd() validation failed: %v", err)
			}
		})
	}
}

func testGenericAppend[T any](t *testing.T, cases []testCaseGenericAppend[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 []T, tt *testCaseGenericAppend[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericAppend() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseGenericAppend[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseGenericAppend[T]) {}
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
			got0 := GenericAppend(
				tt.args.s,
				tt.args.v,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("GenericAppend() validation failed: %v", err)
			}
		})
	}
}

type testGenericAddWants[T constraints.Ordered] struct {
	want0 T
}

type testCaseGenericAdd[T constraints.Ordered] struct {
	name string
	args struct {
		a T
		b T
	}
	want     testGenericAddWants[T]
	init     func(t *testing.T, tt *testCaseGenericAdd[T])
	cleanup  func(t *testing.T, tt *testCaseGenericAdd[T])
	validate func(t *testing.T, got0 T, tt *testCaseGenericAdd[T]) error
}

type testGenericAppendWants[T any] struct {
	want0 []T
}

type testCaseGenericAppend[T any] struct {
	name string
	args struct {
		s []T
		v T
	}
	want     testGenericAppendWants[T]
	init     func(t *testing.T, tt *testCaseGenericAppend[T])
	cleanup  func(t *testing.T, tt *testCaseGenericAppend[T])
	validate func(t *testing.T, got0 []T, tt *testCaseGenericAppend[T]) error
}
