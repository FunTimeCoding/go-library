package argument

func RequiredIntegerFlag(name string) int {
	return RequiredIntegerFlagExit(name, 1)
}
