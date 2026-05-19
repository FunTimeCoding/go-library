package sweep

import "os"

func synchronize(
	source string,
	destination string,
) action {
	di, e := os.Stat(destination)

	if e != nil {
		return actionCopy
	}

	si, f := os.Stat(source)

	if f != nil {
		return actionSkip
	}

	if si.Size() != di.Size() || si.ModTime().After(di.ModTime()) {
		return actionUpdate
	}

	return actionSkip
}
