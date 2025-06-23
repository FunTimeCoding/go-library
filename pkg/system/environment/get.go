package environment

func Get(name string) string {
	return GetExit(name, 1)
}
