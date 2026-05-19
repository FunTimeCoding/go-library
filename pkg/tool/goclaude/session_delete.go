package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionDelete(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id-or-name>",
		Short: "Delete a session and its data",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			_, e := c.DeleteSessionByIdWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			fmt.Printf("deleted: %s\n", identifier)
		},
	}
}
