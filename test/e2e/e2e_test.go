package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	// 1. Compile the binary locally to avoid /tmp noexec issues
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	binPath := filepath.Join(cwd, "gotest2-e2e")

	// Ensure we clean up the binary
	t.Cleanup(func() {
		os.Remove(binPath)
	})

	// Assuming running from project root or logic to find it
	cmd := exec.Command("go", "build", "-o", binPath, "../../cmd/gotest2")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v", err)
	}

	// 2. Iterate over scenarios in testdata
	testdataDir := "testdata"
	entries, err := os.ReadDir(testdataDir)
	if err != nil {
		t.Fatalf("failed to read testdata: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}

		name := entry.Name()
		runTest(t, binPath, testdataDir, name)
	}
}

func runTest(t *testing.T, binPath, srcDir, filename string) {
	t.Run(filename, func(t *testing.T) {
		// Use a separate temp dir for each test to ensure clean state
		tmpDir := t.TempDir()

		// Copy source file
		srcFile := filepath.Join(srcDir, filename)
		srcContent, err := os.ReadFile(srcFile)
		if err != nil {
			t.Fatalf("failed to read source: %v", err)
		}
		destFile := filepath.Join(tmpDir, filename)
		if err := os.WriteFile(destFile, srcContent, 0644); err != nil {
			t.Fatalf("failed to write tmpsource: %v", err)
		}

		// Run gotest2
		cmd := exec.Command(binPath, "-all", destFile)
		// Capture output for debugging
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			t.Fatalf("gotest2 failed: %v\nStderr: %s", err, stderr.String())
		}

		// Verify output file exists
		baseName := strings.TrimSuffix(filename, ".go")
		outTestFile := filepath.Join(tmpDir, baseName+"_test.go")

		if _, err := os.Stat(outTestFile); os.IsNotExist(err) {
			t.Fatalf("output file %s was not created", outTestFile)
		}

		// Read generated content
		got, err := os.ReadFile(outTestFile)
		if err != nil {
			t.Fatalf("failed to read generated file: %v", err)
		}

		// Compare with golden file
		goldenFile := filepath.Join(srcDir, baseName+"_test.golden")
		want, err := os.ReadFile(goldenFile)
		if err != nil {
			t.Fatalf("failed to read golden file %s: %v", goldenFile, err)
		}

		if !bytes.Equal(got, want) {
			t.Errorf("generated content does not match golden file.\nGot:\n%s\nWant:\n%s", got, want)
		}

		// Verify compilation
		// Skip compilation for build_tags.go as it requires specific tags
		if filename == "build_tags.go" {
			return
		}

		// We need to create a go.mod in the tmp dir to make it compiled as a module,
		// or just use run with the file. 'go test -c' needs package context usually.
		// Easiest is to write a dummy go.mod
		if err := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module test\n\ngo 1.20"), 0644); err != nil {
			t.Fatalf("failed to write go.mod: %v", err)
		}

		// Check if it compiles.
		// Note: verification of compilation might fail if the generated test needs imports that are not available
		// (e.g. 3rd party libs) unless we mock them.
		// Our testdata currently only uses stdlib, so it should be fine.
		checkCmd := exec.Command("go", "test", "-c", ".")
		checkCmd.Dir = tmpDir
		if out, err := checkCmd.CombinedOutput(); err != nil {
			t.Errorf("generated code failed to compile: %v\nOutput: %s", err, out)
		}
	})
}
