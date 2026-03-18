package lint

import (
	"fmt"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
	"strings"
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

	var generatedPaths []string

	for _, p := range v.Files() {
		if filepath.Base(p) == library.GeneratedFile {
			generatedPaths = append(generatedPaths, p)
		}
	}

	missing := missingTestDirectories(paths, generatedPaths)
	directories := make([]string, 0, len(missing))

	for d := range missing {
		directories = append(directories, d)
	}

	sort.Strings(directories)

	for _, d := range directories {
		fmt.Printf("%s: %s\n", constant.MissingTestFileText, d)
		name := packageNameOf(v, missing[d])
		fixes.Write(
			filepath.Join(d, stubTestSuffix(name)+"_test.go"),
			stubTestContent(name, strings.Contains(d, "/testdata/")),
		)
	}

	for _, name := range missingSentryPrograms(v) {
		fmt.Printf("%s: cmd/%s\n", constant.MissingSentryText, name)
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
