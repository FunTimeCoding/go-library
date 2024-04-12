package system

func Subdirectories(root string) []string {
	var result []string

	for _, file := range ReadDirectory(root) {
		if file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result
}
