package generator

import (
	"fmt"
	"reflect"
	"testing"
	"text/template"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestFuncMap(t *testing.T) {
	t.Parallel()
	type args struct {
	}
	type wants struct {
		want0 template.FuncMap
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 template.FuncMap, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 template.FuncMap, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("FuncMap() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "exists",
			args: args{},
			want: wants{
				want0: template.FuncMap{
					"add":          add,
					"receiverName": receiverName,
					"isFuzzable":   isFuzzable,
					"isFunc":       isFunc,
					"testName":     getTestFuncName,
					"displayName":  getDisplayFuncName,
				},
			},
			validate: func(t *testing.T, got0 template.FuncMap, tt *test) error {
				if len(got0) != 6 {
					return fmt.Errorf("FuncMap() length = %v, want 6", len(got0))
				}
				for k := range tt.want.want0 {
					if _, ok := got0[k]; !ok {
						return fmt.Errorf("FuncMap() missing key %v", k)
					}
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
			got0 := FuncMap()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("FuncMap() validation failed: %v", err)
			}
		})
	}
}

func Test_add(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	type wants struct {
		want0 int
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 int, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("add() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "1+1",
			args: args{a: 1, b: 1},
			want: wants{want0: 2},
		},
		{
			name: "negative",
			args: args{a: -1, b: -1},
			want: wants{want0: -2},
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
			got0 := add(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("add() validation failed: %v", err)
			}
		})
	}
}

func Test_receiverName(t *testing.T) {
	t.Parallel()
	type args struct {
		t string
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("receiverName() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "pointer",
			args: args{t: "*MyType"},
			want: wants{want0: "MyType"},
		},
		{
			name: "value",
			args: args{t: "MyType"},
			want: wants{want0: "MyType"},
		},
		{
			name: "generic",
			args: args{t: "*List[int]"},
			want: wants{want0: "List"},
		},
		{
			name: "package qualified",
			args: args{t: "model.User"},
			want: wants{want0: "User"},
		},
		{
			name: "unexported to exported",
			args: args{t: "*calculator"},
			want: wants{want0: "Calculator"},
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
			got0 := receiverName(
				tt.args.t,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("receiverName() validation failed: %v", err)
			}
		})
	}
}

func Test_isFuzzable(t *testing.T) {
	t.Parallel()
	type args struct {
		t          string
		typeParams []*models.Field
	}
	type wants struct {
		want0 bool
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 bool, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 bool, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("isFuzzable() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "int",
			args: args{t: "int", typeParams: nil},
			want: wants{want0: true},
		},
		{
			name: "string",
			args: args{t: "string", typeParams: nil},
			want: wants{want0: true},
		},
		{
			name: "unsupported struct",
			args: args{t: "MyStruct", typeParams: nil},
			want: wants{want0: false},
		},
		{
			name: "type param match",
			args: args{
				t:          "T",
				typeParams: []*models.Field{{Name: "T"}},
			},
			want: wants{want0: true},
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
			got0 := isFuzzable(
				tt.args.t,
				tt.args.typeParams,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("isFuzzable() validation failed: %v", err)
			}
		})
	}
}

func Test_getDisplayFuncName(t *testing.T) {
	t.Parallel()
	type args struct {
		fn *models.FunctionInfo
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("getDisplayFuncName() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "simple function",
			args: args{
				fn: &models.FunctionInfo{
					Name: "MyFunc",
				},
			},
			want: wants{want0: "MyFunc"},
		},
		{
			name: "method with pointer receiver",
			args: args{
				fn: &models.FunctionInfo{
					Name: "MyMethod",
					Receiver: &models.Receiver{
						Type: "*MyStruct",
					},
				},
			},
			want: wants{want0: "MyStruct_MyMethod"},
		},
		{
			name: "method with value receiver",
			args: args{
				fn: &models.FunctionInfo{
					Name: "MyMethod",
					Receiver: &models.Receiver{
						Type: "MyStruct",
					},
				},
			},
			want: wants{want0: "MyStruct_MyMethod"},
		},
		{
			name: "generic method",
			args: args{
				fn: &models.FunctionInfo{
					Name: "MyMethod",
					Receiver: &models.Receiver{
						Type: "*List[int]",
					},
				},
			},
			want: wants{want0: "List_MyMethod"},
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
			got0 := getDisplayFuncName(
				tt.args.fn,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("getDisplayFuncName() validation failed: %v", err)
			}
		})
	}
}

func Test_isFunc(t *testing.T) {
	t.Parallel()
	type args struct {
		t string
	}
	type wants struct {
		want0 bool
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 bool, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 bool, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("isFunc() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "func type",
			args: args{t: "func()"},
			want: wants{want0: true},
		},
		{
			name: "non-func type",
			args: args{t: "int"},
			want: wants{want0: false},
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
			got0 := isFunc(
				tt.args.t,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("isFunc() validation failed: %v", err)
			}
		})
	}
}
