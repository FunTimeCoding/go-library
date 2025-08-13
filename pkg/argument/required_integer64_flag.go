package argument

func RequiredInteger64Flag(name string) int64 {
	return RequiredInteger64FlagExit(name, 1)
}
