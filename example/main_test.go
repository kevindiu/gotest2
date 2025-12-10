package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	t.Parallel()
	type args struct {
	}
	type test struct {
		name     string
		args     args
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, tt *test) error
	}
	defaultValidate := func(t *testing.T, tt *test) error {
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
			main()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("_main() validation failed: %v", err)
			}
		})
	}
}
