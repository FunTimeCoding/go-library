package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
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
	fixes := Check(v, skip, verbose)

	if fix {
		fixes.Flush(constant.CurrentDirectory)
	}
}
