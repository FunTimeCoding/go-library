package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionShow(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "show <id-or-name>",
		Short: "Show session detail with alias, description, completions, and summary",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			response, e := c.GetSessionDetailWithResponse(
				context.Background(),
				arguments[0],
			)
			errors.PanicOnError(e)

			if response.JSON200 == nil {
				fmt.Printf("session not found: %s\n", arguments[0])

				return
			}

			d := response.JSON200
			fmt.Printf("Identifier: %s\n", d.Identifier)

			if d.Alias != nil {
				fmt.Printf("Alias: %s\n", *d.Alias)
			}

			if d.Slug != nil {
				fmt.Printf("Slug: %s\n", *d.Slug)
			}

			if d.Created != nil {
				fmt.Printf("Created: %s\n", *d.Created)
			}

			if d.TurnCount != nil {
				fmt.Printf("Turns: %d\n", *d.TurnCount)
			}

			if d.Description != nil {
				fmt.Printf("\n%s\n", *d.Description)
			}

			if d.Completions != nil && len(*d.Completions) > 0 {
				fmt.Println("\nCompletions:")

				for _, o := range *d.Completions {
					fmt.Printf("  [%s] %s\n", o.Kind, o.Topic)

					if o.Summary != nil {
						fmt.Printf("    %s\n", *o.Summary)
					}
				}
			}

			if d.Summary != nil {
				fmt.Printf("\nSummary:\n%s\n", *d.Summary)
			}
		},
	}
}
