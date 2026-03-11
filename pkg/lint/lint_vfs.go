package lint

import (
	"fmt"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/vfs"
	"strings"
)

func LintVFS(
	v *vfs.VFS,
	skip *option.Lint,
	verbose bool,
) *vfs.VFS {
	fixes := vfs.New()
	paths := goFiles(v, skip, verbose)

	read := func(p string) string {
		if fixes.Has(p) {
			return fixes.Read(p)
		}

		return v.Read(p)
	}

	for _, p := range paths {
		if verbose {
			fmt.Printf("Process import: %s\n", p)
		}

		apply(fixes, p, Import(p, strings.NewReader(read(p))))
	}

	for _, p := range paths {
		if verbose {
			fmt.Printf("Process function: %s\n", p)
		}

		apply(fixes, p, Function(p, strings.NewReader(read(p))))
	}

	for _, p := range paths {
		if verbose {
			fmt.Printf("Process variable: %s\n", p)
		}

		apply(fixes, p, Variable(p, strings.NewReader(read(p))))
	}

	for _, p := range paths {
		if verbose {
			fmt.Printf("Process package name: %s\n", p)
		}

		apply(fixes, p, PackageName(p, strings.NewReader(read(p))))
	}

	for _, p := range paths {
		if verbose {
			fmt.Printf("Process type count: %s\n", p)
		}

		apply(fixes, p, TypeCount(p, strings.NewReader(read(p))))
	}

	for _, dir := range missingTestDirs(paths) {
		fmt.Printf("%s: %s\n", lintConstant.MissingTestFileText, dir)
	}

	for _, p := range markupFiles(v, skip, verbose) {
		if verbose {
			fmt.Printf("Process markup: %s\n", p)
		}

		apply(fixes, p, Markup(p, strings.NewReader(read(p))))
	}

	return fixes
}
