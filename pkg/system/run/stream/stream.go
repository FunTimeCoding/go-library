package stream

import (
	"io"
	"os/exec"
)

type Stream struct {
	reader  io.ReadCloser
	command *exec.Cmd
}
