package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionRename(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "rename <id-or-name> <new-name>",
		Short: "Set a display name for a session",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			name := arguments[1]
			_, e := c.Client().PostAliasWithResponse(
				context.Background(),
				client.PostAliasJSONRequestBody{
					Session: identifier,
					Name:    &name,
				},
			)
			errors.PanicOnError(e)
			fmt.Printf("renamed: %s → %s\n", identifier, arguments[1])
		},
	}
}
