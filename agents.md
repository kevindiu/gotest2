# Agent Documentation

This file provides context for AI agents working on `gotest2`.

## Project Overview
`gotest2` is a Go test generator that supports modern Go features, specifically generics. It is an alternative to `gotests`.

## Architecture
- **Parser**: `internal/parser` uses `golang.org/x/tools/go/packages` to parse Go code. Key structs: `FunctionInfo`, `Receiver`, `Field`.
- **Generator**: `internal/generator` uses `text/template` and `embed` to generate test files.
- **Templates**: `internal/generator/templates` contains broken-down templates:
    - `header.tmpl`: Package info.
    - `standard.tmpl`: For non-generic functions.
    - `generic.tmpl`: For generic functions (wrapper strategy).
    - `body.tmpl`: Entry point.

## Generic Strategy
We use a **Wrapper Strategy** for generics:
1. `TestXxx` (Caller): Defines test cases and loops over them. Calls the helper. Contains concrete types (e.g., `int` as default).
2. `testXxx` (Helper): Generic function that implementation the assertion logic.

## Key Files
- `cmd/gotest2/main.go`: CLI entry point.
- `internal/generator/generator.go`: Core generation logic.
- `internal/parser/parser.go`: Parsing logic.

## Testing
Run tests with:
```bash
go test -v ./...
```
Integration tests are in `example/`.

## Workflows
We have defined standard workflows in `.agent/workflows/` to automate common tasks:
- **Verify Changes**: `.agent/workflows/verify.md`
- **Regenerate Examples**: `.agent/workflows/regenerate.md`
- **Add New Feature**: `.agent/workflows/new_feature.md`

## Verification Steps
To verify any new changes:
1. **Regenerate Example Tests**:
   ```bash
   go run cmd/gotest2/main.go ./example/...
   ```
2. **Run Tests**:
   ```bash
   go test -v ./example
   ```
3. **Implement Manual Test Case**:
   Open the generated `_test.go` file (e.g., `example/standard_test.go`) and REPLACE the `// TODO: Add test cases` block with a REAL test case that exercises the feature you are testing (e.g., verifying `Init` hook execution or asserting return values).
   *Run the test again* to confirm your logic passes.
4. **Verify Parallel Execution**:
   Ensure `t.Parallel()` is present in generated files and tests pass.
5. **Dogfooding**:
   Generate tests for internal packages to ensure self-hosting capability.
   ```bash
   go run cmd/gotest2/main.go -- internal/parser/parser.go
   go test -v ./internal/parser
   ```

## Advanced Features

### Lifecycle Hooks
Generated test cases now support `Init` and `Cleanup` hooks:
- `Init(t *testing.T, tt *TestCase)`: Runs before the function call.
- `Cleanup(t *testing.T, tt *TestCase)`: Runs after the function call (via defer).
Useful for setting up/tearing down mocks or fixtures.

### Custom Validation
To override the default `reflect.DeepEqual` check, provide a `Validate` function in your test case:
- Signature: `func(t *testing.T, got1 T1, got2 T2..., tt *TestCase) error`
- If `Validate` fails (returns error), the test fails.
- If `Validate` succeeds (returns nil), default checks are skipped.

### Fuzzing
Generate fuzz targets using the `--fuzz` flag:
```bash
go run cmd/gotest2/main.go --fuzz -- example/fuzz.go
```
This generates `FuzzXxx` functions for any fuzz-compatible inputs.

### Custom Templates
Override default templates:
```bash
go run cmd/gotest2/main.go --template custom.tmpl -- example/standard.go
```

### Verifying Generic Tests with Multiple Types

Generated tests for generic functions include a **Generic Runner** (`runTestXxx[T any]`) and a main test function (`TestXxx`). You can easily add test cases for different concrete types by adding `t.Run` blocks in the main test function.

**Example `generics_test.go`**:
```go
func TestGenericSum(t *testing.T) {
    // ...
    // Integer cases
    t.Run("int", func(t *testing.T) {
        runTestGenericSum[int](t, []testGenericSumTestCase[int]{
            {
                name: "Positive int",
                args: struct{ a, b int }{1, 2},
                want: testGenericSumWants[int]{want0: 3},
            },
        })
    })

    // Float cases
    t.Run("float64", func(t *testing.T) {
        runTestGenericSum[float64](t, []testGenericSumTestCase[float64]{
             {
                name: "Mixed float",
                args: struct{ a, b float64 }{1.5, 2.5},
                want: testGenericSumWants[float64]{want0: 4.0},
            },
        })
    })
}
```
This pattern ensures your generic logic works correctly across all intended types.
