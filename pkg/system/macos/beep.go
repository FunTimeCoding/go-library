package macos

import (
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Beep() {
	errors.PanicOnError(mack.Beep(1))
}
