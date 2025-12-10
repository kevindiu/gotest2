package parser

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestParse(t *testing.T) {
	t.Parallel()
	type args struct {
		patterns []string
	}
	type wants struct {
		want0   []*models.FileResult
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Parse() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("Parse() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "single file",
			init: func(t *testing.T, tt *test) {
				tmpDir := t.TempDir()
				code := `package main
func Foo() {}
`
				filePath := filepath.Join(tmpDir, "main.go")
				if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
					t.Fatal(err)
				}
				tt.args.patterns = []string{filePath}
			},
			validate: func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if len(got0) != 1 {
					return fmt.Errorf("want 1 result, got %d", len(got0))
				}
				if len(got0[0].Functions) != 1 {
					return fmt.Errorf("want 1 function, got %d", len(got0[0].Functions))
				}
				if got0[0].Functions[0].Name != "Foo" {
					return fmt.Errorf("want Func Foo, got %s", got0[0].Functions[0].Name)
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
			got0, err := Parse(
				tt.args.patterns,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("Parse() validation failed: %v", err)
			}
		})
	}
}

func Test_processFunction(t *testing.T) {
	t.Parallel()
	type args struct {
		funcObj    *types.Func
		funcs      *[]*models.FunctionInfo
		importsMap map[string]struct{}
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
			processFunction(
				tt.args.funcObj,
				tt.args.funcs,
				tt.args.importsMap,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("processFunction() validation failed: %v", err)
			}
		})
	}
}

func Test_extractTypeParams(t *testing.T) {
	t.Parallel()
	type args struct {
		sig       *types.Signature
		info      *models.FunctionInfo
		qualifier types.Qualifier
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
			extractTypeParams(
				tt.args.sig,
				tt.args.info,
				tt.args.qualifier,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("extractTypeParams() validation failed: %v", err)
			}
		})
	}
}

func Test_extractReceiver(t *testing.T) {
	t.Parallel()
	type args struct {
		sig       *types.Signature
		info      *models.FunctionInfo
		qualifier types.Qualifier
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
			extractReceiver(
				tt.args.sig,
				tt.args.info,
				tt.args.qualifier,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("extractReceiver() validation failed: %v", err)
			}
		})
	}
}

func Test_extractParams(t *testing.T) {
	t.Parallel()
	type args struct {
		sig       *types.Signature
		info      *models.FunctionInfo
		qualifier types.Qualifier
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
			extractParams(
				tt.args.sig,
				tt.args.info,
				tt.args.qualifier,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("extractParams() validation failed: %v", err)
			}
		})
	}
}

func Test_extractResults(t *testing.T) {
	t.Parallel()
	type args struct {
		sig       *types.Signature
		info      *models.FunctionInfo
		qualifier types.Qualifier
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
			extractResults(
				tt.args.sig,
				tt.args.info,
				tt.args.qualifier,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("extractResults() validation failed: %v", err)
			}
		})
	}
}

func TestParseTests(t *testing.T) {
	t.Parallel()
	type args struct {
		path string
	}
	type wants struct {
		want0   map[string]string
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("ParseTests() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("ParseTests() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "existing test",
			init: func(t *testing.T, tt *test) {
				tmpDir := t.TempDir()
				code := `package main
import "testing"
func TestExisting(t *testing.T) {}
`
				filePath := filepath.Join(tmpDir, "main_test.go")
				if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
					t.Fatal(err)
				}
				tt.args.path = filePath
				tt.want.want0 = map[string]string{
					"TestExisting": "func TestExisting(t *testing.T) {}",
				}
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
			got0, err := ParseTests(
				tt.args.path,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("ParseTests() validation failed: %v", err)
			}
		})
	}
}
