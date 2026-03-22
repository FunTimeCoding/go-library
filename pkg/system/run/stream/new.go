package stream

import (
	"io"
	"os/exec"
)

func New(
	reader io.ReadCloser,
	c *exec.Cmd,
) *Stream {
	return &Stream{reader: reader, command: c}
}
