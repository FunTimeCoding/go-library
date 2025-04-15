package strings

func Swap(
	v []string,
	posA int,
	posB int,
) {
	swap := v[posA]
	v[posA] = v[posB]
	v[posB] = swap
}
