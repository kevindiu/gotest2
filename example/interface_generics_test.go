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
		runTestBox_Get[int](t, []testBox_GetTestCase[int]{
			{
				name:     "Get int",
				receiver: &Box[int]{val: 42},
				want:     testBox_GetWants[int]{want0: 42},
			},
		})
	})
}

func TestBox_Put(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestBox_Put[int](t, []testBox_PutTestCase[int]{
			{
				name:     "Put int",
				receiver: &Box[int]{},
				args: struct {
					v int
				}{v: 100},
				Validate: func(t *testing.T, tt *testBox_PutTestCase[int]) error {
					if tt.receiver.val != 100 {
						return fmt.Errorf("expected 100, got %d", tt.receiver.val)
					}
					return nil
				},
			},
		})
	})
}

func TestProcessContainer(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestProcessContainer[int](t, []testProcessContainerTestCase[int]{
			{
				name: "Process Box",
				args: struct {
					c   Container[int]
					val int
				}{
					c:   &Box[int]{},
					val: 55,
				},
				Validate: func(t *testing.T, tt *testProcessContainerTestCase[int]) error {
					b, ok := tt.args.c.(*Box[int])
					if !ok {
						return fmt.Errorf("expected Box[int]")
					}
					if b.val != 55 {
						return fmt.Errorf("expected 55, got %d", b.val)
					}
					return nil
				},
			},
		})
	})
}

func runTestBox_Get[T any](t *testing.T, cases []testBox_GetTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, tt *testBox_GetTestCase[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("Box_Get() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := tt.receiver.Get()
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("Box_Get() validation failed: %v", err)
			}
		})
	}
}

func runTestBox_Put[T any](t *testing.T, cases []testBox_PutTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, tt *testBox_PutTestCase[T]) error {
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
			tt.receiver.Put(
				tt.args.v,
			)

			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, &tt); err != nil {
				t.Errorf("Box_Put() validation failed: %v", err)
			}
		})
	}
}

func runTestProcessContainer[T any](t *testing.T, cases []testProcessContainerTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, tt *testProcessContainerTestCase[T]) error {
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
			ProcessContainer(
				tt.args.c,
				tt.args.val,
			)

			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, &tt); err != nil {
				t.Errorf("ProcessContainer() validation failed: %v", err)
			}
		})
	}
}

type testBox_GetWants[T any] struct {
	want0 T
}

type testBox_GetTestCase[T any] struct {
	name     string
	receiver *Box[T]
	args     struct {
	}
	want testBox_GetWants[T]

	Init     func(t *testing.T, tt *testBox_GetTestCase[T])
	Cleanup  func(t *testing.T, tt *testBox_GetTestCase[T])
	Validate func(t *testing.T, got0 T, tt *testBox_GetTestCase[T]) error
}

type testBox_PutTestCase[T any] struct {
	name     string
	receiver *Box[T]
	args     struct {
		v T
	}

	Init     func(t *testing.T, tt *testBox_PutTestCase[T])
	Cleanup  func(t *testing.T, tt *testBox_PutTestCase[T])
	Validate func(t *testing.T, tt *testBox_PutTestCase[T]) error
}

type testProcessContainerTestCase[T any] struct {
	name string
	args struct {
		c   Container[T]
		val T
	}

	Init     func(t *testing.T, tt *testProcessContainerTestCase[T])
	Cleanup  func(t *testing.T, tt *testProcessContainerTestCase[T])
	Validate func(t *testing.T, tt *testProcessContainerTestCase[T]) error
}
