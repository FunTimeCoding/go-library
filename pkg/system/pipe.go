package system

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func Pipe(
	c *exec.Cmd,
	input string,
) (string, string, int) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	pipe, pipeFail := c.StdinPipe()
	errors.PanicOnError(pipeFail)
	errors.PanicOnError(c.Start())

	written, writeFail := pipe.Write([]byte(input))
	errors.PanicOnError(writeFail)
	errors.PanicOnError(pipe.Close())

	if f := c.Wait(); f != nil {
		fmt.Printf("Wait fail: %e\n", f)
	}

	return stdout.String(), stderr.String(), written
}
