package issue

func Keys(s string) []string {
	return KeyMatch.FindAllString(s, -1)
}
