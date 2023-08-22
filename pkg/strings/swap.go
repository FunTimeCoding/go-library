package strings

func Swap(
	elements []string,
	posA int,
	posB int,
) {
	swap := elements[posA]
	elements[posA] = elements[posB]
	elements[posB] = swap
}
