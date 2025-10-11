package system

import "os"

func DirectoryExists(path string) bool {
	r, e := os.Stat(path)

	if os.IsNotExist(e) {
		return false
	}

	return r.IsDir()
}
