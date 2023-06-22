package system

import "os"

func FileExists(path string) bool {
	result, e := os.Stat(path)

	if os.IsNotExist(e) {
		return false
	}

	return result.IsDir() == false
}
