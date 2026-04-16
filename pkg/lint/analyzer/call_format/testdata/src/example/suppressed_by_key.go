package example

func suppressedByKey() {
	twoArgs(
		"something-long-enough",
		"to-push-this-line-well-past-the-eighty-character-column-limit-easily",
	) // goanalyze:ignore call_format
}
