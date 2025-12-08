package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Reverse() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0 := Reverse(
				tt.args.s,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("Reverse() validation failed: %v", err)
			}
		})
	}
}
