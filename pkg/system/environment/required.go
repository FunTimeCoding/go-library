package environment

func Required(name string) string {
	return RequiredExit(name, 1)
}
