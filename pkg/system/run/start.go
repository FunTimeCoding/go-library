package run

import (
	"bytes"
	"errors"
	"fmt"
	library "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
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
		r.Print()
	}

	var f *exec.ExitError

	if errors.As(e, &f) {
		r.Exit = f.ExitCode()
	}

	if r.Panic {
		library.PanicOnError(e)
	}

	return r.OutputString
}

func (r *Run) ParseNotation(a any) {
	notation.DecodeStrict(r.OutputString, &a, r.Verbose)
}
