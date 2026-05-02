package gofix

func runVariableNamingFix(
	patterns []string,
	diff bool,
	r *results,
) {
	runVariableNamingFixWithDirectory(patterns, "", diff, r)
}
