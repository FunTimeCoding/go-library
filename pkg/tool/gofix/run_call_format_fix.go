package gofix

func runCallFormatFix(
	patterns []string,
	diff bool,
) {
	runCallFormatFixWithDirectory(patterns, "")
}
