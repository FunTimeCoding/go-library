package strings

func Swap(
	v []string,
	a int,
	b int,
) {
	swap := v[a]
	v[a] = v[b]
	v[b] = swap
}
