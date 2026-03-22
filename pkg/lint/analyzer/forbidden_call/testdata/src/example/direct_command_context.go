package example

import (
	"context"
	"os/exec"
)

func DirectCommandContext() {
	exec.CommandContext( // want `use pkg/system/run instead of exec\.CommandContext`
		context.Background(),
		"ls",
	)
}
