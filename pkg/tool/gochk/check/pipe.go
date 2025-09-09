package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"os/exec"
)

func Pipe(
	c *exec.Cmd,
	input string,
	verbose bool,
) string {
	stdout, stderr := system.Pipe(c, input)

	if verbose {
		if stdout != "" {
			fmt.Printf("Pipe output: %s\n", stdout)
		}

		if stderr != "" {
			fmt.Printf("Pipe error: %s\n", stderr)
		}
	}

	return stdout
}
