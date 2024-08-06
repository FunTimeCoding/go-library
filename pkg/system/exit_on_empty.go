package system

func ExitOnEmpty(
	code int,
	s string,
	text string,
) {
	if s == "" {
		Exitf(code, "empty: %s\n", text)
	}
}
