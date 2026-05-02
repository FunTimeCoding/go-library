package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"os"
)

func Lint(
	skipString string,
	verbose bool,
	fix bool,
	summary bool,
) {
	skip := option.New(skipString, verbose)
	r := output.NewResults()

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

		if fix {
			system.Remove(p)
			r.Add(p, "removed empty directory")
		} else {
			r.AddBlocked(p, "empty directory")
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

		if fix {
			system.Remove(p)
			r.Add(p, "removed empty file")
		} else {
			r.AddBlocked(p, "empty file")
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
				r.Add(p, "removed stray binary")
			} else {
				r.AddBlocked(p, "stray binary")
			}
		}
	}

	fixes := Check(v, skip, fix, verbose, false, &r)

	if fix {
		fixes.Flush(constant.CurrentDirectory)
	}

	hasBlocked := output.PrintResults(r.Entries, summary)

	if hasBlocked {
		os.Exit(1)
	}
}
