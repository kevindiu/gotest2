package app

import (
	"fmt"
	"os"
	"path/filepath"
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
		{
			name: "run simple",
			init: func(t *testing.T, tt *test) {
				tmpDir := t.TempDir()
				code := `package main
func Foo() {}
`
				srcPath := filepath.Join(tmpDir, "main.go")
				if err := os.WriteFile(srcPath, []byte(code), 0644); err != nil {
					t.Fatal(err)
				}
				tt.args.patterns = []string{srcPath}
				tt.args.cfg = Config{
					Tests: true,
				}
			},
			validate: func(t *testing.T, gotErr error, tt *test) error {
				if gotErr != nil {
					return gotErr
				}
				// Check that main_test.go exists
				srcPath := tt.args.patterns[0]
				base := srcPath[:len(srcPath)-3]
				testPath := base + "_test.go"
				if _, err := os.Stat(testPath); os.IsNotExist(err) {
					return fmt.Errorf("test file %s not created", testPath)
				}
				return nil
			},
		},
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
