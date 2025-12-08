---
description: Verify changes by running all tests and validating example generation
---

1. **Run Unit Tests**:
   Run all tests in the project (including internal packages and examples).
   ```bash
   go test -v ./...
   ```

2. **Regenerate Examples**:
   Ensure that regenerating the example tests doesn't break anything or produce unexpected diffs.
   ```bash
   // turbo
   go run cmd/gotest2/main.go ./example/...
   ```

3. **Verify Examples Compile and Pass**:
   After regeneration, run the example tests specifically.
   ```bash
   go test -v ./example/...
   ```

4. **Specific Verification**:
   If you modified creating or validation logic, ensure `complex_test.go` or `edge_cases_test.go` still pass with your changes.
   ```bash
   go test -v -run TestParseConfig ./example/...
   ```
