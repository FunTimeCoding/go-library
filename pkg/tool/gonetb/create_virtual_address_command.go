package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createVirtualAddressCommand(c *netb.Client) *cobra.Command {
	var interfaceName string
	command := &cobra.Command{
		Use:   "create-virtual-address [vm] [address]",
		Short: "Assign an IP address to a virtual machine interface",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.CreateVirtualAddress(arguments[0], interfaceName, arguments[1]))
		},
	}
	command.Flags().StringVar(
		&interfaceName,
		"interface",
		"",
		"interface name (required)",
	)
	errors.PanicOnError(command.MarkFlagRequired("interface"))

	return command
}
