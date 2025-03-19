package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/discord"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const DeleteLoopArgument = "delete-loop"

func main() {
	pflag.String(
		DeleteLoopArgument,
		"",
		"Delete messages in a loop, requires channel",
	)
	argument.ParseAndBind()

	c := discord.NewEnvironment()
	c.Listen(false)
	defer c.Close()

	if h := viper.GetString(DeleteLoopArgument); h != "" {
		c.DeleteLoop(h)

		return
	}

	if false {
		c.PrintStatus()
	}

	system.KillSignalBlock()
}
