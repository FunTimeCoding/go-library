package run

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
)

func (r *Run) Pipe(input string, s ...string) (string, string) {
	c := r.build(s...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	pipe, e := c.StdinPipe()
	errors.PanicOnError(e)
	errors.PanicOnError(c.Start())
	written, e := pipe.Write([]byte(input))
	errors.PanicOnError(e)
	errors.PanicOnError(pipe.Close())

	if f := c.Wait(); f != nil {
		log.Panicf("wait: %e", f)
	}

	if written == 0 {
		log.Panic("no bytes written")
	}

	return stdout.String(), stderr.String()
}
