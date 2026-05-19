package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionPeek(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "peek <id-or-name>",
		Short: "Scan a session for naming",
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

			response, e := c.GetSessionPeekWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			p := response.JSON200
			fmt.Printf(
				"%d lines, %d user messages\n\n",
				p.LineCount,
				len(p.UserMessages),
			)
			limit := len(p.UserMessages) / 4

			if limit < 10 {
				limit = 10
			}

			if limit > len(p.UserMessages) {
				limit = len(p.UserMessages)
			}

			for _, m := range p.UserMessages[:limit] {
				fmt.Println(m)
			}
		},
	}
}
