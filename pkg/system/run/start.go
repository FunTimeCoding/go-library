package run

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func (r *Run) Start(s ...string) string {
	c := exec.Command(s[0], s[1:]...)

	if r.Directory != "" {
		c.Dir = r.Directory
	}

	if len(r.Environment) > 0 {
		c.Env = r.Environment
	}

	result, e := c.CombinedOutput()
	errors.PanicOnError(e)

	return string(result)
}
