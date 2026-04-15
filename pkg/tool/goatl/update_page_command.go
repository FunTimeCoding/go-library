package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func updatePageCommand(c *atl.Client) *cobra.Command {
	var message string
	command := &cobra.Command{
		Use:   "update-page [identifier] [title] [body]",
		Short: "Update a Confluence page with markdown content",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var m *string

			if message != "" {
				m = &message
			}

			fmt.Println(c.UpdatePage(arguments[0], arguments[1], arguments[2], m))
		},
	}
	command.Flags().StringVar(&message, "message", "", "version comment")

	return command
}
