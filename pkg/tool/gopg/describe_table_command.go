package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/client"
	"github.com/spf13/cobra"
)

func describeTableCommand(c *client.Client) *cobra.Command {
	var instance string
	var schema string
	command := &cobra.Command{
		Use:   "describe-table [table]",
		Short: "Show columns, types, and constraints for a table",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			p := &client.DescribeTableParams{Instance: instance}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := c.DescribeTable(
				context.Background(),
				arguments[0],
				p,
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
	command.Flags().StringVar(
		&schema,
		"schema",
		"",
		"Schema name (default: public)",
	)
	errors.PanicOnError(command.MarkFlagRequired("instance"))

	return command
}
