package sweep

import (
	"os"
	"path/filepath"
)

func SyncFile(
	source string,
	harbor string,
) bool {
	destination := filepath.Join(harbor, filepath.Base(source))
	di, e := os.Stat(destination)

	if e != nil {
		copyFile(source, destination)

		return false
	}

	si, f := os.Stat(source)

	if f != nil {
		return false
	}

	if si.Size() == di.Size() && !si.ModTime().After(di.ModTime()) {
		return false
	}

	if si.Size() > di.Size() {
		appendDelta(source, destination, di.Size())

		return false
	}

	copyFile(source, destination)

	return true
}
