package example

import "os/exec"

func DirectCommand() {
	exec.Command("ls") // want `use pkg/system/run instead of exec\.Command`
}
