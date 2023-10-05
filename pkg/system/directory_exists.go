package system

import "os"

func DirectoryExists(path string) bool {
	result, e := os.Stat(path)

	if os.IsNotExist(e) {
		return false
	}

	return result.IsDir()
}
