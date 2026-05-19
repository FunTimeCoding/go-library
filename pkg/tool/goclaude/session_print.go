package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionPrint(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "print <id-or-name>",
		Short: "Print a session to stdout",
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

			response, e := c.GetSessionMessagesWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)

			for _, m := range response.JSON200.Messages {
				if m.Text == "" {
					continue
				}

				if m.IsMeta != nil && *m.IsMeta {
					continue
				}

				fmt.Printf("## %s\n\n%s\n\n", m.Role, m.Text)
			}
		},
	}
}
