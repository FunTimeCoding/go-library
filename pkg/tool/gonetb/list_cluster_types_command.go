package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listClusterTypesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cluster-types",
		Short: "List all NetBox cluster types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListClusterTypes())
		},
	}
}
