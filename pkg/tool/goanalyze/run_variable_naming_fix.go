package goanalyze

func runVariableNamingFix(
	patterns []string,
	diff bool,
) {
	runVariableNamingFixWithDirectory(patterns, "", diff)
}
