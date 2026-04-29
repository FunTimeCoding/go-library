package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/client"
	"github.com/spf13/cobra"
)

func listIndexesCommand(c *client.Client) *cobra.Command {
	var instance string
	var schema string
	var table string
	command := &cobra.Command{
		Use:   "list-indexes",
		Short: "List indexes on a table",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListIndexesParams{
				Instance: instance,
				Table:    table,
			}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := c.ListIndexes(context.Background(), p)
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
		&table,
		"table",
		"",
		"Table name (required)",
	)
	command.Flags().StringVar(
		&schema,
		"schema",
		"",
		"Schema name (default: public)",
	)
	errors.PanicOnError(command.MarkFlagRequired("instance"))
	errors.PanicOnError(command.MarkFlagRequired("table"))

	return command
}
