package gofix

import "os"

func fileExists(path string) bool {
	_, e := os.Stat(path)

	return e == nil
}
