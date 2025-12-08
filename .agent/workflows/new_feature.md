---
description: Guide for adding a new feature to gotest2
---

1. **Update Parser (if needed)**:
   - Modify `internal/models/models.go` if you need new data fields (e.g., tags, comments).
   - Modify `internal/parser/parser.go` to extract this information from `go/packages`.
   - Update `internal/parser/parser_test.go` to verify extraction.

2. **Update Templates**:
   - Modify `templates/` files (e.g., `standard.tmpl`, `generic.tmpl`) to use the new data.
   - Run `go generate` or ensure `embed` picks up changes (usually automatic).

3. **Update Generator**:
   - If logic is complex, update `internal/generator/generator.go` to prepare data before passing to templates.

4. **Verify with Examples**:
   - Add a case to `example/` that uses the new feature (e.g., a new function in `example/standard.go`).
   - Run regeneration: `go run cmd/gotest2/main.go ./example/...`
   - Inspect the generated `_test.go` file manually to ensure it looks correct.
   - Run the test: `go test -v ./example/...`
