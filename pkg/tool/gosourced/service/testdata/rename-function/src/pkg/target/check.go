package target

func Check(content string) bool {
	return IsGeneratedHeader(content)
}
