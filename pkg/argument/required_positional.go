package argument

func RequiredPositional(
	n int,
	name string,
) string {
	return RequiredPositionalExit(n, name, 1)
}
