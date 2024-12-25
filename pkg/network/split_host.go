package network

func SplitHost(s string) string {
	result, _ := SplitHostPort(s)

	return result
}
