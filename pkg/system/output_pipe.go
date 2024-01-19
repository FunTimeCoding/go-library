package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"os/exec"
)

func OutputPipe(c *exec.Cmd) io.ReadCloser {
	result, e := c.StdoutPipe()
	errors.PanicOnError(e)

	return result
}
