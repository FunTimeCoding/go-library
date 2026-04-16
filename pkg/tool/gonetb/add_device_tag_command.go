package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func addDeviceTagCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-device-tag [device] [tag]",
		Short: "Add a tag to a device",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.AddDeviceTag(arguments[0], arguments[1]))
		},
	}
}
