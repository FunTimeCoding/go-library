package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	pattern := filepath.Join(".", "*.go")
	files := slice.StripSuffix(system.Glob(pattern), "_test.go")

	if len(files) == 0 {
		fmt.Println("No relevant files")

		return
	}

	packageRegex := regexp.MustCompile(`^package\s+\w+`)
	importRegex := regexp.MustCompile(`^import\s+`)
	importBlockRegex := regexp.MustCompile(`^import\s*\(`)
	funcRegex := regexp.MustCompile(`^(func|type|var|const)\s+`)

	var packageLine string
	importSet := make(map[string]bool)
	var codeLines []string

	for _, file := range files {
		lines := split.NewLine(system.ReadFile(file))
		inImportBlock := false
		skipEmptyLines := true

		for _, line := range lines {
			trimmed := strings.TrimSpace(line)

			// Handle package declaration - only keep the first one
			if packageRegex.MatchString(trimmed) {
				if packageLine == "" {
					packageLine = line
				}

				continue
			}

			// Handle import blocks and single imports
			if importBlockRegex.MatchString(trimmed) {
				inImportBlock = true

				continue
			} else if inImportBlock && trimmed == ")" {
				inImportBlock = false
				skipEmptyLines = true // Skip empty lines after imports

				continue
			} else if inImportBlock {
				if trimmed != "" {
					importSet[trimmed] = true
				}

				continue
			} else if importRegex.MatchString(trimmed) {
				importSet[trimmed[7:]] = true
				skipEmptyLines = true

				continue
			}

			// Skip empty lines immediately after imports/package
			if trimmed == "" && skipEmptyLines {
				continue
			}

			// Once we hit actual code, stop skipping empty lines
			if trimmed != "" {
				skipEmptyLines = false
			}

			// Add the line to our code collection
			codeLines = append(codeLines, line)
		}
	}

	// Output package and imports
	if packageLine != "" {
		fmt.Println(packageLine)
		fmt.Println()
	}

	if len(importSet) > 0 {
		var imports []string

		for imp := range importSet {
			imports = append(imports, imp)
		}

		sort.Strings(imports)
		fmt.Println("import (")

		for _, imp := range imports {
			fmt.Printf("\t%s\n", imp)
		}

		fmt.Println(")")
		fmt.Println()
	}

	// Normalize spacing in code section
	prevWasEmpty := false

	for i, line := range codeLines {
		trimmed := strings.TrimSpace(line)
		isEmpty := trimmed == ""
		isTopLevel := funcRegex.MatchString(trimmed)

		// Add blank line before top-level declarations (except the very first one)
		if isTopLevel && i > 0 && !prevWasEmpty {
			fmt.Println()
		}

		// Print the line (skip consecutive empty lines)
		if !isEmpty || !prevWasEmpty {
			fmt.Println(line)
		}

		prevWasEmpty = isEmpty
	}
}
