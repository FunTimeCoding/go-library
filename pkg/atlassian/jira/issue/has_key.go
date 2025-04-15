package issue

func HasKey(s string) bool {
	return KeyMatch.MatchString(s)
}
