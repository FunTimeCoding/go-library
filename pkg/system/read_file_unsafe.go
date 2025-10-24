package system

func ReadFileUnsafe(name string) string {
	return string(ReadBytesUnsafe(name))
}
