package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"os/exec"
	"runtime"
)

func OpenBrowser(locator string) {
	switch p := runtime.GOOS; p {
	case Linux:
		errors.PanicOnError(exec.Command("xdg-open", locator).Start())
	case Darwin:
		errors.PanicOnError(exec.Command("open", locator).Start())
	default:
		unexpected.String(p)
	}
}
