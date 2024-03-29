package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func LookPath(name string) string {
	result, e := exec.LookPath(name)
	errors.PanicOnError(e)

	return result
}
