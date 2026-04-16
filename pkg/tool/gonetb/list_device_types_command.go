package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listDeviceTypesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-types",
		Short: "List all NetBox device types",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(c.ListDeviceTypes())
		},
	}
}
