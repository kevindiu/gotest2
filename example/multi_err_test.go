package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiError(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	type wants struct {
		want0    int
		wantErr1 error
		wantErr2 error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 int, gotErr1 error, gotErr2 error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, gotErr1 error, gotErr2 error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("MultiError() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr1) != fmt.Sprint(tt.want.wantErr1) {
			return fmt.Errorf("MultiError() error1 = %v, wantErr1 %v", gotErr1, tt.want.wantErr1)
		}
		if fmt.Sprint(gotErr2) != fmt.Sprint(tt.want.wantErr2) {
			return fmt.Errorf("MultiError() error2 = %v, wantErr2 %v", gotErr2, tt.want.wantErr2)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
			got0, err1, err2 := MultiError(
				tt.args.n,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err1, err2, &tt); err != nil {
				t.Errorf("MultiError() validation failed: %v", err)
			}
		})
	}
}
