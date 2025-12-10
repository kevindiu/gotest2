package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenericMultiError(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testGenericMultiError[int](t, []testCaseGenericMultiError[int]{
			// TODO: Add test cases.
		})
	})
}

func testGenericMultiError[T any](t *testing.T, cases []testCaseGenericMultiError[T]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, gotErr1 error, gotErr2 error, tt *testCaseGenericMultiError[T]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("GenericMultiError() got0 = %v, want %v", got0, tt.want.want0)
			}
			if fmt.Sprint(gotErr1) != fmt.Sprint(tt.want.wantErr1) {
				return fmt.Errorf("GenericMultiError() error1 = %v, wantErr1 %v", gotErr1, tt.want.wantErr1)
			}
			if fmt.Sprint(gotErr2) != fmt.Sprint(tt.want.wantErr2) {
				return fmt.Errorf("GenericMultiError() error2 = %v, wantErr2 %v", gotErr2, tt.want.wantErr2)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseGenericMultiError[T]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseGenericMultiError[T]) {}
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
			got0, err1, err2 := GenericMultiError[T](
				tt.args.val,
				tt.args.n,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err1, err2, &tt); err != nil {
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

type testCaseGenericMultiError[T any] struct {
	name string
	args struct {
		val T
		n   int
	}
	want     testGenericMultiErrorWants[T]
	init     func(t *testing.T, tt *testCaseGenericMultiError[T])
	cleanup  func(t *testing.T, tt *testCaseGenericMultiError[T])
	validate func(t *testing.T, got0 T, gotErr1 error, gotErr2 error, tt *testCaseGenericMultiError[T]) error
}
