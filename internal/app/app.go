package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kevindiu/gotest2/internal/generator"
	"github.com/kevindiu/gotest2/internal/models"
	"github.com/kevindiu/gotest2/internal/parser"
)

// Config holds the application configuration.
type Config struct {
	All          bool
	Exported     bool
	TemplatePath string
	Parallel     bool
	Fuzz         bool
	Benchmark    bool
	Tests        bool
}

// Run executes the application logic.
func Run(patterns []string, cfg Config) error {
	fmt.Printf("Processing %d patterns: %v\n", len(patterns), patterns)
	results, err := parser.Parse(patterns)
	if err != nil {
		return fmt.Errorf("failed to parse packages: %v", err)
	}

	for _, result := range results {
		fmt.Printf("Processing %s (%s)...\n", result.Path, result.PackageName)

		if len(result.Functions) == 0 {
			fmt.Printf("No functions found to generate tests for in %s\n", result.Path)
			continue
		}

		// Filter based on flags
		var targetFuncs []*models.FunctionInfo
		for _, fn := range result.Functions {
			if cfg.All {
				targetFuncs = append(targetFuncs, fn)
				continue
			}
			if cfg.Exported && !fn.IsExported {
				continue
			}
			targetFuncs = append(targetFuncs, fn)
		}

		if len(targetFuncs) == 0 {
			continue
		}

		// Define generation configurations
		type GenConfig struct {
			Enabled     bool
			Suffix      string
			Template    string
			RunParallel bool
			RunFuzz     bool
			RunBench    bool
			RunTests    bool
		}

		configs := []GenConfig{
			{
				Enabled:     cfg.Tests,
				Suffix:      "_test.go",
				Template:    "body.tmpl",
				RunParallel: cfg.Parallel,
				RunFuzz:     false,
				RunBench:    false,
				RunTests:    true,
			},
			{
				Enabled:     cfg.Benchmark,
				Suffix:      "_bench_test.go",
				Template:    "body_bench.tmpl",
				RunParallel: false,
				RunFuzz:     false,
				RunBench:    true,
				RunTests:    false,
			},
			{
				Enabled:     cfg.Fuzz,
				Suffix:      "_fuzz_test.go",
				Template:    "body_fuzz.tmpl",
				RunParallel: false,
				RunFuzz:     true,
				RunBench:    false,
				RunTests:    false,
			},
		}

		for _, c := range configs {
			if !c.Enabled {
				continue
			}

			// Determine output path
			absPath := result.Path
			ext := filepath.Ext(absPath)
			base := absPath[:len(absPath)-len(ext)]
			testPath := base + c.Suffix

			// Parse existing tests
			var existingTests map[string]string
			if _, err := os.Stat(testPath); err == nil {
				existingTests, err = parser.ParseTests(testPath)
				if err != nil {
					log.Printf("Warning: failed to parse existing tests in %s: %v", testPath, err)
				}
			}

			// Generate code
			code, err := generator.Generate(targetFuncs, result.Imports, result.PackageName, cfg.TemplatePath, c.RunParallel, c.RunFuzz, c.RunBench, c.RunTests, existingTests, c.Template)
			if err != nil {
				log.Printf("Failed to generate code for %s: %v", testPath, err)
				if len(code) > 0 {
					debugPath := testPath + ".debug"
					os.WriteFile(debugPath, code, 0644)
					log.Printf("Wrote unformatted code to %s", debugPath)
				}
				continue
			}

			// Only write if there is content (or if we want to update file)
			if len(code) > 0 {
				if err := os.WriteFile(testPath, code, 0644); err != nil {
					log.Printf("Failed to write to %s: %v", testPath, err)
				} else {
					fmt.Printf("Generated %s\n", testPath)
				}
			}
		}
	}
	return nil
}
