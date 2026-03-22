package example

import (
	"context"
	"os/exec"
)

func DirectCommandContext() {
	exec.CommandContext(
		context.Background(),
		"ls",
	) // want `use pkg/system/run instead of exec\.CommandContext`
}
