package gofix

func runVariableNamingFix(
	patterns []string,
	diff bool,
) {
	runVariableNamingFixWithDirectory(patterns, "", diff)
}
