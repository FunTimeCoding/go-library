package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Go(
	main string,
	output string,
) {
	var name string

	if main == "" {
		main = "."
	} else {
		name = filepath.Base(filepath.Dir(main))
		fmt.Printf("Name: %s\n", name)
	}

	if output == "" {
		if name == "" {
			panic("output empty and main not specified")
		} else {
			output = system.Join(system.Temporary, name)
		}
	}

	fmt.Printf("Output: %s\n", output)
	path := filepath.Dir(output)
	fmt.Printf("Path: %s\n", path)
	system.EnsurePathExists(path)
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
	fmt.Printf("Command: %s\n", join.Space(s))
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
