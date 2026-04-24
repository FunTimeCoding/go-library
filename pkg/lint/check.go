package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
)

func Check(
	v *virtual_file_system.System,
	skip *option.Lint,
	fix bool,
	verbose bool,
	stubTest bool,
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
			StrayConstant,
			Spacing,
			VariableGrouping,
		},
		fix,
		verbose,
	)

	if stubTest {
		generateStubTests(v, fixes, paths, fix)
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
