package system

import "io/fs"

func Mode(path string) fs.FileMode {
	return Stat(path).Mode()
}
