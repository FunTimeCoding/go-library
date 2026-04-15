package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listSitesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-sites",
		Short: "List all NetBox sites",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListSites())
		},
	}
}
