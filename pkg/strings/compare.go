package strings

func Compare(
	past []string,
	now []string,
) ([]string, []string, []string) {
	add := difference(now, past)
	remove := difference(past, now)
	remain := common(past, now)

	return add, remove, remain
}
