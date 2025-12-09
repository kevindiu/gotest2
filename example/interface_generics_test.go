package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBox_Get(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testBox_Get[int](t, []testCaseBox_Get[int]{
			// TODO: Add test cases.
		})
	})
}

func TestBox_Put(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testBox_Put[int](t, []testCaseBox_Put[int]{
			// TODO: Add test cases.
		})
	})
}

func TestProcessContainer(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testProcessContainer[int](t, []testCaseProcessContainer[int]{
			// TODO: Add test cases.
		})
	})
}

func testBox_Get[T any](t *testing.T, cases []testCaseBox_Get[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testCaseBox_Get[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("Box_Get() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseBox_Get[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseBox_Get[T]) {}
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
			got0 := tt.receiver.Get()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("Box_Get() validation failed: %v", err)
			}
		})
	}
}

func testBox_Put[T any](t *testing.T, cases []testCaseBox_Put[T]) {
	t.Parallel()
	for _, tt := range cases {

		defaultValidate := func(t *testing.T, tt *testCaseBox_Put[T]) error {
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseBox_Put[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseBox_Put[T]) {}
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
			tt.receiver.Put(
				tt.args.v,
			)

			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("Box_Put() validation failed: %v", err)
			}
		})
	}
}

func testProcessContainer[T any](t *testing.T, cases []testCaseProcessContainer[T]) {
	t.Parallel()
	for _, tt := range cases {

		defaultValidate := func(t *testing.T, tt *testCaseProcessContainer[T]) error {
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseProcessContainer[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseProcessContainer[T]) {}
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
			ProcessContainer(
				tt.args.c,
				tt.args.val,
			)

			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("ProcessContainer() validation failed: %v", err)
			}
		})
	}
}

type testBox_GetWants[T any] struct {
	want0 T
}

type testCaseBox_Get[T any] struct {
	name     string
	receiver *Box[T]
	args     struct {
	}
	want     testBox_GetWants[T]
	init     func(t *testing.T, tt *testCaseBox_Get[T])
	cleanup  func(t *testing.T, tt *testCaseBox_Get[T])
	validate func(t *testing.T, got0 T, tt *testCaseBox_Get[T]) error
}

type testCaseBox_Put[T any] struct {
	name     string
	receiver *Box[T]
	args     struct {
		v T
	}
	init     func(t *testing.T, tt *testCaseBox_Put[T])
	cleanup  func(t *testing.T, tt *testCaseBox_Put[T])
	validate func(t *testing.T, tt *testCaseBox_Put[T]) error
}

type testCaseProcessContainer[T any] struct {
	name string
	args struct {
		c   Container[T]
		val T
	}
	init     func(t *testing.T, tt *testCaseProcessContainer[T])
	cleanup  func(t *testing.T, tt *testCaseProcessContainer[T])
	validate func(t *testing.T, tt *testCaseProcessContainer[T]) error
}
