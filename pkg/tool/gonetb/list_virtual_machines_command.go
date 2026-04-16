package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listVirtualMachinesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-machines",
		Short: "List all NetBox virtual machines",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListVirtualMachines())
		},
	}
}
