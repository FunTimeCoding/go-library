package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createDeviceTypeCommand(c *netb.Client) *cobra.Command {
	var manufacturer string
	command := &cobra.Command{
		Use:   "create-device-type [model]",
		Short: "Create a NetBox device type",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.CreateDeviceType(arguments[0], manufacturer))
		},
	}
	command.Flags().StringVar(
		&manufacturer,
		"manufacturer",
		"",
		"manufacturer name (required)",
	)
	errors.PanicOnError(command.MarkFlagRequired("manufacturer"))

	return command
}
