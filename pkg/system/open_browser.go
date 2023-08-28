package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
	"runtime"
)

func OpenBrowser(locator string) {
	var e error

	switch p := runtime.GOOS; p {
	case Linux:
		e = exec.Command("xdg-open", locator).Start()
	case Darwin:
		e = exec.Command("open", locator).Start()
	default:
		e = fmt.Errorf("unexpected platform: %s", p)
	}

	errors.PanicOnError(e)
}
