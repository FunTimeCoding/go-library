package argument

func Required(name string) string {
	return RequiredExit(name, 1)
}
