package system

func ReadFile(
	base string,
	name string,
) string {
	return string(ReadBytes(base, name))
}
