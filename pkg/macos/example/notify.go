package example

import (
	"github.com/funtimecoding/go-library/pkg/macos"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func Notify() {
	macos.Notify(strings.Alfa, strings.Bravo, strings.Charlie)
	macos.SimpleNotify(strings.Alfa)
	macos.Beep()
	macos.Alert("Subject", "Body")
	macos.InputDialog("Test1", "Test2", "")
	macos.CustomDialog("Test1", "Test2")
}
