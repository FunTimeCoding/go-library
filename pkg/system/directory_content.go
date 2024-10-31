package system

func DirectoryContent(path string) []string {
	var result []string

	for _, element := range ReadDirectory(path) {
		result = append(result, element.Name())
	}

	return result
}
