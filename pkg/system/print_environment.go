package system

import (
	"fmt"
	"os"
	"strings"
)

func PrintEnvironment() {
	for _, e := range os.Environ() {
		p := strings.SplitN(e, "=", 2)
		fmt.Printf("ENV: %s=%s\n", p[0], p[1])
	}
}
