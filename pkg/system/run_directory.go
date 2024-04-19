package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func RunDirectory(
	directory string,
	s ...string,
) string {
	c := exec.Command(s[0], s[1:]...)
	c.Dir = directory
	result, e := c.CombinedOutput()
	errors.PanicOnError(e)

	return string(result)
}
