package runtime

import (
	"os"
	"strings"
)

func RunningFromBinary() bool {
	return !strings.Contains(os.Args[0], "go-build")
}
