package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listAddressesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-addresses [device]",
		Short: "List IP addresses for a NetBox device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.ListAddresses(arguments[0]))
		},
	}
}
