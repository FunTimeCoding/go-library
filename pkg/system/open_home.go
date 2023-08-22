package system

import "os"

func OpenHome(subPath string) *os.File {
	return Open(Join(Home(), subPath))
}
