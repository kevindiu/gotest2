package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGenericMultiError(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		runTestGenericMultiError[int](t, []testGenericMultiErrorTestCase[int]{
			{
				name: "Negative input",
				args: struct {
					val int
					n   int
				}{val: 10, n: -1},
				want: testGenericMultiErrorWants[int]{
					want0:    10,
					wantErr1: errors.New("negative"),
				},
			},
			{
				name: "Too large input",
				args: struct {
					val int
					n   int
				}{val: 10, n: 101},
				want: testGenericMultiErrorWants[int]{
					want0:    0,
					wantErr2: errors.New("too large"),
				},
			},
			{
				name: "Normal input",
				args: struct {
					val int
					n   int
				}{val: 42, n: 50},
				want: testGenericMultiErrorWants[int]{
					want0: 42,
				},
			},
		})
	})
	t.Run("string", func(t *testing.T) {
		runTestGenericMultiError[string](t, []testGenericMultiErrorTestCase[string]{
			{
				name: "Negative input string",
				args: struct {
					val string
					n   int
				}{val: "foo", n: -1},
				want: testGenericMultiErrorWants[string]{
					want0:    "foo",
					wantErr1: errors.New("negative"),
				},
			},
			{
				name: "Too large input string",
				args: struct {
					val string
					n   int
				}{val: "foo", n: 101},
				want: testGenericMultiErrorWants[string]{
					want0:    "",
					wantErr2: errors.New("too large"),
				},
			},
			{
				name: "Normal input string",
				args: struct {
					val string
					n   int
				}{val: "success", n: 50},
				want: testGenericMultiErrorWants[string]{
					want0: "success",
				},
			},
		})
	})
}

func runTestGenericMultiError[T any](t *testing.T, cases []testGenericMultiErrorTestCase[T]) {
	t.Parallel()

	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, gotErr1 error, gotErr2 error, tt *testGenericMultiErrorTestCase[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericMultiError() got0 = %v, want %v", got0, tt.want.want0)
			}
			if (gotErr1 != nil) != (tt.want.wantErr1 != nil) {
				return fmt.Errorf("GenericMultiError() error1 = %v, wantErr1 %v", gotErr1, tt.want.wantErr1)
			}
			if gotErr1 != nil && tt.want.wantErr1 != nil && gotErr1.Error() != tt.want.wantErr1.Error() {
				return fmt.Errorf("GenericMultiError() error1 = %v, wantErr1 %v", gotErr1, tt.want.wantErr1)
			}
			if (gotErr2 != nil) != (tt.want.wantErr2 != nil) {
				return fmt.Errorf("GenericMultiError() error2 = %v, wantErr2 %v", gotErr2, tt.want.wantErr2)
			}
			if gotErr2 != nil && tt.want.wantErr2 != nil && gotErr2.Error() != tt.want.wantErr2.Error() {
				return fmt.Errorf("GenericMultiError() error2 = %v, wantErr2 %v", gotErr2, tt.want.wantErr2)
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
			got0, err1, err2 := GenericMultiError(
				tt.args.val,
				tt.args.n,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, err1, err2, &tt); err != nil {
				t.Errorf("GenericMultiError() validation failed: %v", err)
			}
		})
	}
}

type testGenericMultiErrorWants[T any] struct {
	want0    T
	wantErr1 error
	wantErr2 error
}

type testGenericMultiErrorTestCase[T any] struct {
	name string
	args struct {
		val T
		n   int
	}
	want testGenericMultiErrorWants[T]

	Init     func(t *testing.T, tt *testGenericMultiErrorTestCase[T])
	Cleanup  func(t *testing.T, tt *testGenericMultiErrorTestCase[T])
	Validate func(t *testing.T, got0 T, gotErr1 error, gotErr2 error, tt *testGenericMultiErrorTestCase[T]) error
}
