package argument

func RequiredPositional(
	n int,
	description string,
) string {
	return RequiredPositionalExit(n, description, 1)
}
