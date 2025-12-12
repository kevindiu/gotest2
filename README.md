# gotest2 - Modern Go Test Generator

`gotest2` is a CLI tool similar to `gotests` but built with modern Go features in mind, specifically **Generics**. It leverages `golang.org/x/tools/go/packages` to correctly parse and generate tests for generic functions and methods.

## Features

- **Generics Support**: First-class support for Go 1.18+ generics.
    - Uses a **wrapper strategy** where test logic is in a generic helper `testXxx[T]` and data is defined in `TestXxx` with concrete types.
- **Parallel Execution**: Generates parallel-safe tests (`t.Parallel()`) by default. Disable with `--parallel=false`.
- **Fuzzing Support**: Generates fuzz targets (`FuzzXxx`) for suitable functions using `--fuzz`.
- **Lifecycle Hooks**:
    - `Init` and `Cleanup` hooks for per-test-case setup/teardown.
    - **Enhanced Access**: Hooks have access to the test case struct (via `tt`), allowing dynamic initialization of arguments or receivers.
- **Incremental Generation**: Parses existing test files and preserves your custom modifications to `TestXxx` functions while updating runners and structs.
- **Standard Go Support**: Full support for standard functions and methods.
- **Smart Naming**: Auto-generates names like `TestReceiver_Method` to avoid collisions.
- **Custom Templates**: Use your own templates with `--template`.

## Installation

```bash
go install github.com/kevindiu/gotest2@latest
```

## Usage

Generate tests for specific files:

```bash
gotest2 path/to/file.go
```

Generate parallel tests (default) with fuzzing targets:

```bash
gotest2 --fuzz path/to/file.go
```

Use a custom template:

```bash
gotest2 --template my.tmpl path/to/file.go
```

## Comparisons with `gotests`

| Feature | `gotests` | `gotest2` |
| :--- | :--- | :--- |
| **Generics** | Limited/No support | **Full Support** |
| **Parser** | `go/ast` | `go/packages` |
| **Templates** | Customizable | Customizable |
| **Existing Tests** | Preserves | **Preserves / Incremental** |
| **Parallel** | Supported | **Supported Default** |
| **Fuzzing** | No support | **Supported** |
| **Lifecycle** | No explicit support | **Init/Cleanup Hooks** |

## License

MIT
