package run

import (
	"bytes"
	"errors"
	"fmt"
	library "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"os/exec"
)

func (r *Run) Start(s ...string) string {
	if r.Verbose {
		fmt.Printf("Run: %s\n", join.Space(s...))
	}

	c := r.build(s...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
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
