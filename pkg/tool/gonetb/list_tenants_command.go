package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listTenantsCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tenants",
		Short: "List all NetBox tenants",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(c.ListTenants())
		},
	}
}
