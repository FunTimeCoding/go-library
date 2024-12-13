package expression

func SubMatch(
	expression string,
	search string,
) string {
	return SubMatchIndex(expression, search, 1)
}
