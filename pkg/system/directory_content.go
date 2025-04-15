package system

func DirectoryContent(path string) []string {
	var result []string

	for _, e := range ReadDirectory(path) {
		result = append(result, e.Name())
	}

	return result
}
