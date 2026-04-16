package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createVirtualInterfaceCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-virtual-interface [vm] [name]",
		Short: "Create an interface on a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateVirtualInterface(arguments[0], arguments[1]))
		},
	}
}
