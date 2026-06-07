package gofix

import "github.com/funtimecoding/go-library/pkg/lint/output"

func runSingleParameterFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runSingleParameterFixWithDirectory(patterns, "", r)
}
