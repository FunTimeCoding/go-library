package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
)

func Check(
	v *virtual_file_system.System,
	skip *option.Lint,
	fix bool,
	verbose bool,
	stubTest bool,
	r *output.Results,
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
			StrayConstant,
			Spacing,
			VariableGrouping,
		},
		fix,
		verbose,
		r,
	)

	if stubTest {
		generateStubTests(v, fixes, paths, fix)
	}

	for _, name := range missingSentryPrograms(v) {
		r.AddBlocked(
			fmt.Sprintf("cmd/%s", name),
			"missing sentry reporter",
		)
	}

	runCheckers(
		v,
		fixes,
		markupFiles(v, skip, verbose),
		[]Checker{Markup},
		fix,
		verbose,
		r,
	)

	return fixes
}
