package run

import (
	"bytes"
	"errors"
	"fmt"
	errorLibrary "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"os"
	"os/exec"
)

func (r *Run) Start(s ...string) string {
	if r.Verbose {
		fmt.Printf("Run: %s\n", join.Space(s...))
	}

	c := exec.Command(s[0], s[1:]...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if r.Directory != "" {
		c.Dir = r.Directory
	}

	if len(r.environment) > 0 {
		c.Env = append(os.Environ(), r.environment...)
	}

	e := c.Run()
	r.OutputString = stdout.String()
	r.ErrorString = stderr.String()

	if e != nil {
		r.Error = e
	}

	if r.Verbose {
		if r.Error != nil {
			fmt.Printf("Error: %s\n", r.Error)
		}

		if r.OutputString != "" {
			fmt.Printf("Stdout:\n%s\n", r.OutputString)
		}

		if r.ErrorString != "" {
			fmt.Printf("Stderr:\n%s\n", r.ErrorString)
		}
	}

	var f *exec.ExitError

	if errors.As(e, &f) {
		r.ExitCode = f.ExitCode()
	}

	if r.Panic {
		errorLibrary.PanicOnError(e)
	}

	return r.OutputString
}
