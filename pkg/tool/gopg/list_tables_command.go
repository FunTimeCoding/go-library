package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/client"
	"github.com/spf13/cobra"
)

func listTablesCommand(c *client.Client) *cobra.Command {
	var instance string
	var schema string
	command := &cobra.Command{
		Use:   "list-tables",
		Short: "List tables in a schema",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListTablesParams{Instance: instance}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := c.ListTables(context.Background(), p)
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
	command.Flags().StringVar(
		&schema,
		"schema",
		"",
		"Schema name (default: public)",
	)
	errors.PanicOnError(command.MarkFlagRequired("instance"))

	return command
}
