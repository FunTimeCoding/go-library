package argument

func RequiredInteger(name string) int {
	return RequiredIntegerExit(name, 1)
}
