package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listClustersCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-clusters",
		Short: "List all NetBox clusters",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListClusters())
		},
	}
}
