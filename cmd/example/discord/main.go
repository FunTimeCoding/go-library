package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/discord"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	DeleteLoopFlag  = "delete-loop"
	ChannelArgument = "channel"
)

func main() {
	pflag.BoolP(
		DeleteLoopFlag,
		"d",
		false,
		"Delete messages in a loop",
	)
	pflag.StringP(
		ChannelArgument,
		"c",
		"",
		"Channel to delete from",
	)
	argument.ParseAndBind()

	c := discord.NewEnvironment()
	c.Listen(false)
	defer c.Close()

	if viper.GetBool(DeleteLoopFlag) {
		h := viper.GetString(ChannelArgument)

		if h == "" {
			fmt.Println("Channel is required")

			return
		}

		c.DeleteLoop(h)

		return
	}

	if false {
		c.PrintStatus()
	}

	system.KillSignalBlock()
}
