package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Lint(
	skipString string,
	verbose bool,
	fix bool,
) {
	skip := option.New(skipString, verbose)

	for _, p := range system.EmptyDirectories(constant.CurrentDirectory) {
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

	for _, p := range system.FilesRecursive(constant.CurrentDirectory) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip file: %s\n", p)
			}

			continue
		}

		if !system.FileEmpty(p) {
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

		goFiles = append(goFiles, p)
	}

	for _, p := range goFiles {
		f := system.Open(p)
		Import(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range goFiles {
		f := system.Open(p)
		Function(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}

	for _, p := range system.MarkupFiles(constant.CurrentDirectory) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip markup file: %s\n", p)
			}

			continue
		}

		f := system.Open(p)
		Markup(p, f).ProcessConcerns(fix)
		errors.LogClose(f)
	}
}
