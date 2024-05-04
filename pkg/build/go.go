package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"os/exec"
	"runtime"
)

func Go(
	main string,
	output string,
) {
	if main == "" {
		main = "."
	}

	s := []string{
		"go",
		"build",
		"-ldflags",
		fmt.Sprintf(
			"-X main.Version=%s -X main.GitHash=%s -X main.BuildDate=%s",
			GitTag(),
			GitHash(),
			Date(),
		),
		"-o",
		output,
		main,
	}
	c := exec.Command(s[0], s[1:]...)
	c.Env = os.Environ()
	c.Env = append(c.Env, "CGO_ENABLED=0")
	c.Env = append(c.Env, fmt.Sprintf("GOOS=%s", runtime.GOOS))
	combined, e := c.CombinedOutput()

	if t := string(combined); t != "" {
		fmt.Printf("Output:\n%s", t)
	}

	errors.PanicOnError(e)
}
