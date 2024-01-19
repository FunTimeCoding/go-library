package system

import (
	"errors"
	errorHelper "github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func ExitErrorCode(e error) int {
	if e != nil {
		var f *exec.ExitError

		if errors.As(e, &f) {
			return f.ExitCode()
		} else {
			errorHelper.PanicOnError(e)
		}
	}

	return 0
}
