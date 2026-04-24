package argument

func RequiredInteger64(name string) int64 {
	return RequiredInteger64Exit(name, 1)
}
