package gopostgres

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func explain(c *client.Client) *cobra.Command {
	var instance string
	var analyze bool
	command := &cobra.Command{
		Use:   "explain [sql]",
		Short: "Show execution plan for a SQL query",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.Explain(
				context.Background(),
				client.ExplainJSONRequestBody{
					Instance: instance,
					Sql:      arguments[0],
					Analyze:  &analyze,
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
	command.Flags().BoolVar(
		&analyze,
		"analyze",
		false,
		"Run EXPLAIN ANALYZE",
	)
	errors.PanicOnError(command.MarkFlagRequired("instance"))

	return command
}
