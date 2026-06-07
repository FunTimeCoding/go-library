package gofix

import "github.com/funtimecoding/go-library/pkg/lint/output"

func runCallFormatFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runCallFormatFixWithDirectory(patterns, "", r)
}
