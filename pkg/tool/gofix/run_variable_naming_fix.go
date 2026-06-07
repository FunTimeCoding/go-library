package gofix

import "github.com/funtimecoding/go-library/pkg/lint/output"

func runVariableNamingFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runVariableNamingFixWithDirectory(patterns, "", diff, r)
}
