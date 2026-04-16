package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func removeVirtualTagCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-virtual-tag [vm] [tag]",
		Short: "Remove a tag from a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.RemoveVirtualTag(arguments[0], arguments[1]))
		},
	}
}
