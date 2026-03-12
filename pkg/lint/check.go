package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
)

func Check(
	v *virtual_file_system.System,
	skip *option.Lint,
	verbose bool,
) *virtual_file_system.System {
	fixes := virtual_file_system.New()
	paths := goFiles(v, skip, verbose)

	runCheckers(
		v,
		fixes,
		paths,
		[]Checker{
			Import,
			Function,
			Variable,
			PackageName,
			TypeCount,
			Spacing,
		},
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
			filepath.Join(dir, stubTestSuffix(name)+"_test.go"),
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
