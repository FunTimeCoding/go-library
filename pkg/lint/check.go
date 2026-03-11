package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/vfs"
	"path/filepath"
	"sort"
)

func Check(
	v *vfs.VFS,
	skip *option.Lint,
	verbose bool,
) *vfs.VFS {
	fixes := vfs.New()
	paths := goFiles(v, skip, verbose)

	runCheckers(
		v,
		fixes,
		paths,
		[]Checker{Import, Function, Variable, PackageName, TypeCount, Spacing},
		verbose,
	)

	missing := missingTestDirs(paths)
	dirs := make([]string, 0, len(missing))

	for dir := range missing {
		dirs = append(dirs, dir)
	}

	sort.Strings(dirs)

	for _, dir := range dirs {
		fmt.Printf("%s: %s\n", constant.MissingTestFileText, dir)
		name := packageNameOf(v, missing[dir])
		fixes.Write(
			filepath.Join(dir, "stub_test.go"),
			stubTestContent(name),
		)
	}

	runCheckers(
		v,
		fixes,
		markupFiles(v, skip, verbose),
		[]Checker{Markup},
		verbose,
	)

	return fixes
}
