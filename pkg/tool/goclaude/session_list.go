package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionList(c *client.ClientWithResponses) *cobra.Command {
	var peek bool
	result := &cobra.Command{
		Use:   "list",
		Short: "List all sessions",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			response, e := c.GetSessionsWithResponse(
				context.Background(),
				&client.GetSessionsParams{Peek: &peek},
			)
			errors.PanicOnError(e)

			for _, s := range response.JSON200.Sessions {
				ts := formatTimestamp(s.Timestamp)
				alias := ""

				if s.Alias != nil {
					alias = *s.Alias
				}

				slug := ""

				if s.Slug != nil {
					slug = *s.Slug
				}

				name := slug
				marker := " "

				if alias != "" {
					name = alias
					marker = "*"
				}

				fmt.Printf(
					"%s  %s  %5d  %s %s",
					ts,
					s.Identifier,
					s.Lines,
					marker,
					name,
				)

				if peek && s.Preview != nil {
					fmt.Printf("  %s", *s.Preview)
				}

				fmt.Println()
			}
		},
	}
	result.Flags().BoolVar(
		&peek,
		"peek",
		false,
		"Show first user message preview",
	)

	return result
}
