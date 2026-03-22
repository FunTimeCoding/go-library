package process

import "os/exec"

type Process struct {
	command *exec.Cmd
}
