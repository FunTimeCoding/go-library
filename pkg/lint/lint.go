package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
	"strings"
)

func Lint(
	skipString string,
	verbose bool,
	fix bool,
) {
	skip := option.New(skipString, verbose)

	for _, p := range system.EmptyDirectories(
		constant.CurrentDirectory,
		verbose,
	) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip empty directory: %s\n", p)
			}

			continue
		}

		fmt.Printf("Empty directory: %s\n", p)

		if fix {
			system.Remove(p)
		}
	}

	for _, p := range system.FilesRecursive(
		constant.CurrentDirectory,
		verbose,
	) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip file: %s\n", p)
			}

			continue
		}

		if !system.FileEmpty(p) {
			if verbose {
				fmt.Printf("Non empty file: %s\n", p)
			}

			continue
		}

		fmt.Printf("Empty file: %s\n", p)

		if fix {
			system.Remove(p)
		}
	}

	var goFiles []string

	for _, p := range system.GoFiles(constant.CurrentDirectory) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip go file: %s\n", p)
			}

			continue
		}

		if verbose {
			fmt.Printf("Select go file: %s\n", p)
		}

		goFiles = append(goFiles, p)
	}

	for _, p := range goFiles {
		if verbose {
			fmt.Printf("Process import: %s\n", p)
		}

		f := system.Open(p)
		Import(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range goFiles {
		if verbose {
			fmt.Printf("Process function: %s\n", p)
		}

		f := system.Open(p)
		Function(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range goFiles {
		if verbose {
			fmt.Printf("Process variable: %s\n", p)
		}

		f := system.Open(p)
		Variable(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range goFiles {
		if verbose {
			fmt.Printf("Process package name: %s\n", p)
		}

		f := system.Open(p)
		PackageName(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range goFiles {
		if verbose {
			fmt.Printf("Process type count: %s\n", p)
		}

		f := system.Open(p)
		TypeCount(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	dirs := make(map[string]bool)

	for _, p := range goFiles {
		dirs[filepath.Dir(p)] = false
	}

	for _, p := range goFiles {
		if strings.HasSuffix(p, "_test.go") {
			dirs[filepath.Dir(p)] = true
		}
	}

	for dir, hasTest := range dirs {
		if !hasTest {
			fmt.Printf(
				"%s: %s\n",
				lintConstant.MissingTestFileText,
				dir,
			)
		}
	}

	for _, p := range system.MarkupFiles(constant.CurrentDirectory) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip markup file: %s\n", p)
			}

			continue
		}

		if verbose {
			fmt.Printf("Process markup: %s\n", p)
		}

		f := system.Open(p)
		Markup(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}
}
