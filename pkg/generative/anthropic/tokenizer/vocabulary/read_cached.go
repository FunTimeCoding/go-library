package vocabulary

import "os"

func readCached() ([]byte, error) {
	return os.ReadFile(cachePath())
}
