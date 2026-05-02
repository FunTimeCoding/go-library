package gofix

func runCallFormatFix(
	patterns []string,
	diff bool,
	r *results,
) {
	runCallFormatFixWithDirectory(patterns, "", r)
}
