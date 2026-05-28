package process

import (
	"io"
	"os/exec"
)

type Process struct {
	command *exec.Cmd
	stdout  io.ReadCloser
	stderr  io.ReadCloser
}
