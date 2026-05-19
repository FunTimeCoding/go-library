package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionExport(c *client.ClientWithResponses) *cobra.Command {
	var all bool
	result := &cobra.Command{
		Use:   "export [id-or-name]",
		Short: "Export a session to markdown",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			if all {
				response, e := c.PostSessionsExportWithResponse(
					context.Background(),
				)
				errors.PanicOnError(e)

				for _, path := range response.JSON200.Paths {
					fmt.Printf("exported: %s\n", path)
				}

				return
			}

			if len(arguments) == 0 {
				fmt.Println("provide a session id or use --all")

				return
			}

			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			response, e := c.PostSessionExportWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)

			for _, path := range response.JSON200.Paths {
				fmt.Printf("exported: %s\n", path)
			}
		},
	}
	result.Flags().BoolVar(
		&all,
		"all",
		false,
		"Export all sessions",
	)

	return result
}
