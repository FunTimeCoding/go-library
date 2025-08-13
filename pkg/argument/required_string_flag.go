package argument

func RequiredStringFlag(name string) string {
	return RequiredStringFlagExit(name, 1)
}
