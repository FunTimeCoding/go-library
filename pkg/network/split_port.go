package network

func SplitPort(s string) int {
	_, result := SplitHostPort(s)

	return result
}
