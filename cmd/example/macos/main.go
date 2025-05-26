package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/macos"
	"github.com/funtimecoding/go-library/pkg/sound"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	if true {
		macos.CreateBundle(
			"Example",
			"tmp",
			"tmp/example",
			"tmp/icon.icns",
			"sh.s3n",
			"1.0.0",
		)
	}

	if false {
		fmt.Println(system.Run(sound.Afplay, sound.SosumiPath))
	}

	if false {
		macos.Notify(strings.Alfa, strings.Bravo, strings.Charlie)
		macos.SimpleNotify(strings.Alfa)
		macos.Beep()
		macos.Alert("Subject", "Body")
		macos.InputDialog("Test1", "Test2", "")
		macos.CustomDialog("Test1", "Test2")
	}

	if false {
		fmt.Println(system.Run(sound.Afplay, sound.TinkPath))
	}
}
