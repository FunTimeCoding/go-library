package process

import "os/exec"

func New(c *exec.Cmd) *Process {
	return &Process{command: c}
}
