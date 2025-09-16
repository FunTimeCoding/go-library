package environment

func Exit(name string) string {
	return ExitCode(name, 1)
}
