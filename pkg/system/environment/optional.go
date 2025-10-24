package environment

func Optional(name string) string {
	return Fallback(name, "")
}
