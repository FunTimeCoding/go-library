package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/errors"
	"os/exec"
	"runtime"
)

func OpenBrowser(locator string) {
	var e error
	platform := runtime.GOOS

	switch platform {
	case "linux":
		e = exec.Command("xdg-open", locator).Start()
	case "darwin":
		e = exec.Command("open", locator).Start()
	default:
		e = fmt.Errorf("unexpected platform: %s", platform)
	}

	errors.PanicOnError(e)
}
