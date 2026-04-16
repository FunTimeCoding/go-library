package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listDeviceRolesCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-roles",
		Short: "List all NetBox device roles",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListDeviceRoles())
		},
	}
}
