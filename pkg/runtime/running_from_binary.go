package runtime

import (
	"os"
	"strings"
)

func RunningFromBinary() bool {
	if strings.Contains(os.Args[0], "go-build") {
		return false
	}

	return true
}
