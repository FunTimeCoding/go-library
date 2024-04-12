package run

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func (r *Run) Start(s ...string) string {
	c := exec.Command(s[0], s[1:]...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if r.Directory != "" {
		c.Dir = r.Directory
	}

	if len(r.Environment) > 0 {
		c.Env = r.Environment
	}

	e := c.Run()
	r.OutputString = stdout.String()
	r.ErrorString = stderr.String()

	if e != nil {
		r.Error = e
	}

	if r.Verbose {
		fmt.Printf("Error: %s\n", r.Error)
		fmt.Printf("Stdout:\n%s\n", r.OutputString)
		fmt.Printf("Stderr:\n%s\n", r.ErrorString)
	}

	if r.Panic {
		errors.PanicOnError(e)
	}

	return r.OutputString
}
