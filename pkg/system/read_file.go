package system

func ReadFile(name string) string {
	return string(ReadBytes(name))
}
