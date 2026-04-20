package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/client"
	"github.com/spf13/cobra"
)

func queryCommand(c *client.Client) *cobra.Command {
	var instance string
	command := &cobra.Command{
		Use:   "query [sql]",
		Short: "Execute SQL against a PostgreSQL instance",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.Query(
				context.Background(),
				client.QueryJSONRequestBody{
					Instance: instance,
					Sql:      arguments[0],
				},
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
