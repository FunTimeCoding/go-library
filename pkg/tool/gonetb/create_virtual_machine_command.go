package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createVirtualMachineCommand(c *netb.Client) *cobra.Command {
	var cluster string
	command := &cobra.Command{
		Use:   "create-virtual-machine [name]",
		Short: "Create a NetBox virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.CreateVirtualMachine(arguments[0], cluster))
		},
	}
	command.Flags().StringVar(&cluster, "cluster", "", "cluster name (required)")
	errors.PanicOnError(command.MarkFlagRequired("cluster"))

	return command
}
