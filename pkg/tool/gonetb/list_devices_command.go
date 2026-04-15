package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listDevicesCommand(c *netb.Client) *cobra.Command {
	var query string
	command := &cobra.Command{
		Use:   "list-devices",
		Short: "List NetBox devices",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var q *string

			if query != "" {
				q = &query
			}

			fmt.Println(c.ListDevices(q))
		},
	}
	command.Flags().StringVar(&query, "query", "", "filter by name")

	return command
}
