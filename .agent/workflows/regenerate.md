---
description: Regenerate all test files in the example directory
---

1. **Run Generator**:
   Execute the tool against the `example/` directory recursively.
   ```bash
   // turbo
   go run cmd/gotest2/main.go ./example/...
   ```

2. **Verify Output**:
   Check if the generated files compile and pass tests.
   ```bash
   go test -v ./example/...
   ```
