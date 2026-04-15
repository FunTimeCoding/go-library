package goanalyze

func runCallFormatFix(
	patterns []string,
	diff bool,
) {
	runCallFormatFixWithDirectory(patterns, "")
}
