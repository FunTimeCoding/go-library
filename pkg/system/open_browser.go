package system

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"runtime"
)

func OpenBrowser(locator string) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		run.New().Launch("xdg-open", locator)
	case constant.Darwin:
		run.New().Launch("open", locator)
	default:
		unexpected.String(p)
	}
}
