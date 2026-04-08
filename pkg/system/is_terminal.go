package system

import "os"

func IsTerminal() bool {
	i, e := os.Stdin.Stat()

	if e != nil {
		return false
	}

	return i.Mode()&os.ModeCharDevice != 0
}
