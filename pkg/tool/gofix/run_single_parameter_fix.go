package gofix

func runSingleParameterFix(
	patterns []string,
	diff bool,
	r *results,
) {
	runSingleParameterFixWithDirectory(patterns, "", r)
}
