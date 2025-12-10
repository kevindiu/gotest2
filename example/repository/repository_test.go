package repository

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMemoryRepository_Create(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMemoryRepository_Create[int, int](t, []testCaseMemoryRepository_Create[int, int]{
			// TODO: Add test cases.
		})
	})
}

func TestMemoryRepository_Get(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMemoryRepository_Get[int, int](t, []testCaseMemoryRepository_Get[int, int]{
			// TODO: Add test cases.
		})
	})
}

func TestMemoryRepository_List(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMemoryRepository_List[int, int](t, []testCaseMemoryRepository_List[int, int]{
			// TODO: Add test cases.
		})
	})
}

func TestMemoryRepository_Delete(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testMemoryRepository_Delete[int, int](t, []testCaseMemoryRepository_Delete[int, int]{
			// TODO: Add test cases.
		})
	})
}

func TestNewMemoryRepository(t *testing.T) {
	t.Parallel()
	// TODO: Add test cases.
	t.Run("int", func(t *testing.T) {
		testNewMemoryRepository[int, int](t, []testCaseNewMemoryRepository[int, int]{
			// TODO: Add test cases.
		})
	})
}

func testMemoryRepository_Create[T any, ID comparable](t *testing.T, cases []testCaseMemoryRepository_Create[T, ID]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, gotErr error, tt *testCaseMemoryRepository_Create[T, ID]) error {
			if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
				return fmt.Errorf("MemoryRepository_Create() error = %v, wantErr %v", gotErr, tt.want.wantErr)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMemoryRepository_Create[T, ID]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMemoryRepository_Create[T, ID]) {}
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
			err := tt.receiver.Create(
				tt.args.id,
				tt.args.entity,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, err, &tt); err != nil {
				t.Errorf("MemoryRepository_Create() validation failed: %v", err)
			}
		})
	}
}

func testMemoryRepository_Get[T any, ID comparable](t *testing.T, cases []testCaseMemoryRepository_Get[T, ID]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 T, gotErr error, tt *testCaseMemoryRepository_Get[T, ID]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("MemoryRepository_Get() got0 = %v, want %v", got0, tt.want.want0)
			}
			if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
				return fmt.Errorf("MemoryRepository_Get() error = %v, wantErr %v", gotErr, tt.want.wantErr)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMemoryRepository_Get[T, ID]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMemoryRepository_Get[T, ID]) {}
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
			got0, err := tt.receiver.Get(
				tt.args.id,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("MemoryRepository_Get() validation failed: %v", err)
			}
		})
	}
}

func testMemoryRepository_List[T any, ID comparable](t *testing.T, cases []testCaseMemoryRepository_List[T, ID]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 []T, gotErr error, tt *testCaseMemoryRepository_List[T, ID]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("MemoryRepository_List() got0 = %v, want %v", got0, tt.want.want0)
			}
			if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
				return fmt.Errorf("MemoryRepository_List() error = %v, wantErr %v", gotErr, tt.want.wantErr)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMemoryRepository_List[T, ID]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMemoryRepository_List[T, ID]) {}
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
			got0, err := tt.receiver.List()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("MemoryRepository_List() validation failed: %v", err)
			}
		})
	}
}

func testMemoryRepository_Delete[T any, ID comparable](t *testing.T, cases []testCaseMemoryRepository_Delete[T, ID]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, gotErr error, tt *testCaseMemoryRepository_Delete[T, ID]) error {
			if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
				return fmt.Errorf("MemoryRepository_Delete() error = %v, wantErr %v", gotErr, tt.want.wantErr)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseMemoryRepository_Delete[T, ID]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseMemoryRepository_Delete[T, ID]) {}
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
			err := tt.receiver.Delete(
				tt.args.id,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, err, &tt); err != nil {
				t.Errorf("MemoryRepository_Delete() validation failed: %v", err)
			}
		})
	}
}

func testNewMemoryRepository[T any, ID comparable](t *testing.T, cases []testCaseNewMemoryRepository[T, ID]) {
	t.Parallel()
	for _, tt := range cases {
		defaultValidate := func(t *testing.T, got0 *MemoryRepository[T, ID], tt *testCaseNewMemoryRepository[T, ID]) error {
			if !reflect.DeepEqual(got0, tt.want.want0) {
				return fmt.Errorf("NewMemoryRepository() got0 = %v, want %v", got0, tt.want.want0)
			}
			return nil
		}
		defaultInit := func(t *testing.T, tt *testCaseNewMemoryRepository[T, ID]) {}
		defaultCleanup := func(t *testing.T, tt *testCaseNewMemoryRepository[T, ID]) {}
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
			got0 := NewMemoryRepository[T, ID]()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("NewMemoryRepository() validation failed: %v", err)
			}
		})
	}
}

type testMemoryRepository_CreateWants[T any, ID comparable] struct {
	wantErr error
}

type testCaseMemoryRepository_Create[T any, ID comparable] struct {
	name     string
	receiver *MemoryRepository[T, ID]
	args     struct {
		id     ID
		entity T
	}
	want     testMemoryRepository_CreateWants[T, ID]
	init     func(t *testing.T, tt *testCaseMemoryRepository_Create[T, ID])
	cleanup  func(t *testing.T, tt *testCaseMemoryRepository_Create[T, ID])
	validate func(t *testing.T, gotErr error, tt *testCaseMemoryRepository_Create[T, ID]) error
}

type testMemoryRepository_GetWants[T any, ID comparable] struct {
	want0   T
	wantErr error
}

type testCaseMemoryRepository_Get[T any, ID comparable] struct {
	name     string
	receiver *MemoryRepository[T, ID]
	args     struct {
		id ID
	}
	want     testMemoryRepository_GetWants[T, ID]
	init     func(t *testing.T, tt *testCaseMemoryRepository_Get[T, ID])
	cleanup  func(t *testing.T, tt *testCaseMemoryRepository_Get[T, ID])
	validate func(t *testing.T, got0 T, gotErr error, tt *testCaseMemoryRepository_Get[T, ID]) error
}

type testMemoryRepository_ListWants[T any, ID comparable] struct {
	want0   []T
	wantErr error
}

type testCaseMemoryRepository_List[T any, ID comparable] struct {
	name     string
	receiver *MemoryRepository[T, ID]
	args     struct {
	}
	want     testMemoryRepository_ListWants[T, ID]
	init     func(t *testing.T, tt *testCaseMemoryRepository_List[T, ID])
	cleanup  func(t *testing.T, tt *testCaseMemoryRepository_List[T, ID])
	validate func(t *testing.T, got0 []T, gotErr error, tt *testCaseMemoryRepository_List[T, ID]) error
}

type testMemoryRepository_DeleteWants[T any, ID comparable] struct {
	wantErr error
}

type testCaseMemoryRepository_Delete[T any, ID comparable] struct {
	name     string
	receiver *MemoryRepository[T, ID]
	args     struct {
		id ID
	}
	want     testMemoryRepository_DeleteWants[T, ID]
	init     func(t *testing.T, tt *testCaseMemoryRepository_Delete[T, ID])
	cleanup  func(t *testing.T, tt *testCaseMemoryRepository_Delete[T, ID])
	validate func(t *testing.T, gotErr error, tt *testCaseMemoryRepository_Delete[T, ID]) error
}

type testNewMemoryRepositoryWants[T any, ID comparable] struct {
	want0 *MemoryRepository[T, ID]
}

type testCaseNewMemoryRepository[T any, ID comparable] struct {
	name string
	args struct {
	}
	want     testNewMemoryRepositoryWants[T, ID]
	init     func(t *testing.T, tt *testCaseNewMemoryRepository[T, ID])
	cleanup  func(t *testing.T, tt *testCaseNewMemoryRepository[T, ID])
	validate func(t *testing.T, got0 *MemoryRepository[T, ID], tt *testCaseNewMemoryRepository[T, ID]) error
}
