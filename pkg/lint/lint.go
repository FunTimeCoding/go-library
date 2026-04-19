package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
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

	v := virtual_file_system.From(constant.CurrentDirectory)

	for _, p := range v.Files() {
		if Skipped(skip, p) {
			continue
		}

		if isKnownBinaryExtension(p) {
			continue
		}

		if isExecutable(v.Read(p)) {
			if fix {
				system.Remove(p)
				fmt.Printf("%s (deleted): %s\n", lintConstant.StrayBinaryText, p)
			} else {
				fmt.Printf("%s: %s\n", lintConstant.StrayBinaryText, p)
			}
		}
	}

	fixes := Check(v, skip, fix, verbose)

	if fix {
		fixes.Flush(constant.CurrentDirectory)
	}
}
