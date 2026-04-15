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
	fix bool,
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
			ForbiddenImport,
			Function,
			Variable,
			PackageName,
			TypeCount,
			FunctionCount,
			StrayConstant,
			Spacing,
			VariableGrouping,
		},
		fix,
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
	useAssert := strings.Contains(
		v.Read("go.mod"),
		library.GoLibraryModule,
	)

	for _, d := range directories {
		if fix {
			fmt.Printf(
				"%s: %s (auto-fixed)\n",
				constant.MissingTestFileText,
				d,
			)
		} else {
			fmt.Printf(
				"%s: %s (auto-fixable)\n",
				constant.MissingTestFileText,
				d,
			)
		}

		name := packageNameOf(v, missing[d])
		isToolPackage := strings.HasPrefix(d, "pkg/tool/") &&
			strings.Count(strings.TrimPrefix(d, "pkg/tool/"), "/") == 0
		filename := fmt.Sprintf("%s_test.go", stubTestSuffix(name))

		if isToolPackage {
			filename = "main_test.go"
		}

		fixes.Write(
			filepath.Join(d, filename),
			stubTestContent(
				name,
				strings.Contains(d, "/testdata/"),
				isToolPackage,
				useAssert,
			),
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
		fix,
		verbose,
	)

	return fixes
}
