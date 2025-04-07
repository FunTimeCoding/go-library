package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
	"testing"
)

func TestOutputPipe(t *testing.T) {
	c := exec.Command("echo", "hello")
	p := OutputPipe(c)
	errors.PanicOnError(c.Start())
	assert.String(t, "hello\n", string(ReadAll(p)))
}
