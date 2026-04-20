package gopg

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/client"
	"github.com/spf13/cobra"
)

func listInstancesCommand(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List configured PostgreSQL instances",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.ListInstances(context.Background())
			errors.PanicOnError(e)
			printResponse(r)
		},
	}
}
