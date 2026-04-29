package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/client"
	"github.com/spf13/cobra"
)

func listSchemasCommand(c *client.Client) *cobra.Command {
	var instance string
	command := &cobra.Command{
		Use:   "list-schemas",
		Short: "List schemas in the active database",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.ListSchemas(
				context.Background(),
				&client.ListSchemasParams{Instance: instance},
			)
			errors.PanicOnError(e)
			printResponse(r)
		},
	}
	command.Flags().StringVar(
		&instance,
		"instance",
		"",
		"Instance name (required)",
	)
	errors.PanicOnError(command.MarkFlagRequired("instance"))

	return command
}
