package system

import (
	"github.com/funtimecoding/go-library/errors"
	"os/exec"
)

func Run(s ...string) string {
	result, e := exec.Command(s[0], s[1:]...).CombinedOutput()
	errors.PanicOnError(e)

	return string(result)
}
