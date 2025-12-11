package parser

import (
	"fmt"
	"go/token"
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
		{
			name: "basic function",
			args: args{
				funcObj: types.NewFunc(token.NoPos, nil, "Foo",
					types.NewSignatureType(nil, nil, nil, nil, nil, false),
				),
				funcs:      &[]*models.FunctionInfo{},
				importsMap: make(map[string]struct{}),
			},
			validate: func(t *testing.T, tt *test) error {
				funcs := *tt.args.funcs
				if len(funcs) != 1 {
					return fmt.Errorf("want 1 function, got %d", len(funcs))
				}
				if funcs[0].Name != "Foo" {
					return fmt.Errorf("want Foo, got %s", funcs[0].Name)
				}
				return nil
			},
		},
		{
			name: "function with imports",
			args: args{
				funcObj: types.NewFunc(token.NoPos, nil, "Bar",
					types.NewSignatureType(nil, nil, nil,
						types.NewTuple(
							types.NewVar(token.NoPos, nil, "r",
								types.NewNamed(
									types.NewTypeName(token.NoPos, types.NewPackage("example.com/pkg", "pkg"), "MyType", nil),
									nil, nil,
								),
							),
						),
						nil, false),
				),
				funcs:      &[]*models.FunctionInfo{},
				importsMap: make(map[string]struct{}),
			},
			validate: func(t *testing.T, tt *test) error {
				if _, ok := tt.args.importsMap["example.com/pkg"]; !ok {
					return fmt.Errorf("expected import example.com/pkg")
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
		{
			name: "no type params",
			args: args{
				sig:       types.NewSignatureType(nil, nil, nil, nil, nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.TypeParams) != 0 {
					return fmt.Errorf("want 0 type params, got %d", len(tt.args.info.TypeParams))
				}
				return nil
			},
		},
		{
			name: "type params",
			args: args{
				sig: types.NewSignatureType(nil, nil,
					[]*types.TypeParam{
						types.NewTypeParam(types.NewTypeName(0, nil, "T", nil), types.NewInterfaceType(nil, nil).Complete()),
					},
					nil, nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.TypeParams) != 1 {
					return fmt.Errorf("want 1 type param, got %d", len(tt.args.info.TypeParams))
				}
				if tt.args.info.TypeParams[0].Name != "T" {
					return fmt.Errorf("want type param T, got %s", tt.args.info.TypeParams[0].Name)
				}
				if tt.args.info.TypeParams[0].Type != "interface{}" {
					return fmt.Errorf("want type param interface{}, got %s", tt.args.info.TypeParams[0].Type)
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
		{
			name: "no receiver",
			args: args{
				sig:       types.NewSignatureType(nil, nil, nil, nil, nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if tt.args.info.Receiver != nil {
					return fmt.Errorf("want nil receiver, got %v", tt.args.info.Receiver)
				}
				return nil
			},
		},
		{
			name: "pointer receiver",
			args: args{
				sig: types.NewSignatureType(
					types.NewVar(0, nil, "r", types.NewPointer(
						types.NewNamed(
							types.NewTypeName(0, nil, "MyType", nil),
							types.NewStruct(nil, nil),
							nil,
						),
					)),
					nil, nil, nil, nil, false,
				),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if tt.args.info.Receiver == nil {
					return fmt.Errorf("want receiver, got nil")
				}
				if tt.args.info.Receiver.Name != "r" {
					return fmt.Errorf("want receiver name 'r', got '%s'", tt.args.info.Receiver.Name)
				}
				if tt.args.info.Receiver.Type != "*MyType" {
					return fmt.Errorf("want receiver type '*MyType', got '%s'", tt.args.info.Receiver.Type)
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
		{
			name: "no params",
			args: args{
				sig:       types.NewSignatureType(nil, nil, nil, nil, nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.Parameters) != 0 {
					return fmt.Errorf("want 0 params, got %d", len(tt.args.info.Parameters))
				}
				return nil
			},
		},
		{
			name: "simple params",
			args: args{
				sig: types.NewSignatureType(nil, nil, nil,
					types.NewTuple(
						types.NewVar(0, nil, "a", types.Typ[types.Int]),
						types.NewVar(0, nil, "b", types.Typ[types.String]),
					),
					nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.Parameters) != 2 {
					return fmt.Errorf("want 2 params, got %d", len(tt.args.info.Parameters))
				}
				if tt.args.info.Parameters[0].Name != "a" || tt.args.info.Parameters[0].Type != "int" {
					return fmt.Errorf("param 0 mismatch")
				}
				if tt.args.info.Parameters[1].Name != "b" || tt.args.info.Parameters[1].Type != "string" {
					return fmt.Errorf("param 1 mismatch")
				}
				return nil
			},
		},
		{
			name: "variadic",
			args: args{
				sig: types.NewSignatureType(nil, nil, nil,
					types.NewTuple(
						types.NewVar(0, nil, "args", types.NewSlice(types.Typ[types.Int])),
					),
					nil, true),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.Parameters) != 1 {
					return fmt.Errorf("want 1 param, got %d", len(tt.args.info.Parameters))
				}
				if !tt.args.info.Parameters[0].IsVariadic {
					return fmt.Errorf("expected variadic")
				}
				if tt.args.info.Parameters[0].Type != "[]int" {
					return fmt.Errorf("want []int, got %s", tt.args.info.Parameters[0].Type)
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
		{
			name: "no results",
			args: args{
				sig:       types.NewSignatureType(nil, nil, nil, nil, nil, false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.Results) != 0 {
					return fmt.Errorf("want 0 results, got %d", len(tt.args.info.Results))
				}
				return nil
			},
		},
		{
			name: "results",
			args: args{
				sig: types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(
						types.NewVar(0, nil, "", types.Universe.Lookup("error").Type()),
					),
					false),
				info:      &models.FunctionInfo{},
				qualifier: func(p *types.Package) string { return "" },
			},
			validate: func(t *testing.T, tt *test) error {
				if len(tt.args.info.Results) != 1 {
					return fmt.Errorf("want 1 result, got %d", len(tt.args.info.Results))
				}
				if tt.args.info.Results[0].Type != "error" {
					return fmt.Errorf("want error, got %s", tt.args.info.Results[0].Type)
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
