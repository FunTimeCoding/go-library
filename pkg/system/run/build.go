package run

import (
	"os"
	"os/exec"
)

func (r *Run) build(s ...string) *exec.Cmd {
	c := exec.Command(s[0], s[1:]...) // goanalyze:ignore

	if r.Directory != "" {
		c.Dir = r.Directory
	}

	if len(r.environment) > 0 {
		c.Env = append(os.Environ(), r.environment...)
	}

	return c
}
