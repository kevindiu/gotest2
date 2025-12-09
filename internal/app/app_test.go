package app

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()
	type args struct {
		patterns []string
		cfg      Config
	}
	type wants struct {
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, gotErr error, tt *test) error {
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("Run() error = %v, wantErr %v", gotErr, tt.want.wantErr)
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
			err := Run(
				tt.args.patterns,
				tt.args.cfg,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, err, &tt); err != nil {
				t.Errorf("Run() validation failed: %v", err)
			}
		})
	}
}
