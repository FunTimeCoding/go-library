package example

import (
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"github.com/funtimecoding/go-library/pkg/system/macos"
)

func Notify() {
	macos.Notify(upper.Alfa, upper.Bravo, upper.Charlie)
	macos.SimpleNotify(upper.Alfa)
	macos.Beep()
	macos.Alert("Subject", "Body")
	macos.InputDialog("Test1", "Test2", "")
	macos.CustomDialog("Test1", "Test2")
}
