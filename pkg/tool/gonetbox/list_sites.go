package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listSites(c *client.Client) *cobra.Command {
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
