package system

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
	"os/exec"
)

func Pipe(
	c *exec.Cmd,
	input string,
) (string, string) {
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
		log.Panicf("wait: %e", f)
	}

	if written == 0 {
		log.Panic("no bytes written")
	}

	return stdout.String(), stderr.String()
}
