package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createTenantCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tenant [name]",
		Short: "Create a NetBox tenant",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateTenant(arguments[0]))
		},
	}
}
