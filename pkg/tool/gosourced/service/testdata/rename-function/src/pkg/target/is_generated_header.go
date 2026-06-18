package target

func IsGeneratedHeader(content string) bool {
	return len(content) > 20
}
