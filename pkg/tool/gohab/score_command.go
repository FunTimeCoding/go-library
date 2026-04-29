package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
	"github.com/spf13/cobra"
)

func scoreCommand(c *client.Client) *cobra.Command {
	var direction string
	command := &cobra.Command{
		Use:   "score [identifier]",
		Short: "Score a task",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.Score(arguments[0], direction))
		},
	}
	command.Flags().StringVar(
		&direction,
		"direction",
		"up",
		"Score direction: up or down",
	)

	return command
}
