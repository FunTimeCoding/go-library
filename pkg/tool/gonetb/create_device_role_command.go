package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createDeviceRoleCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-device-role [name]",
		Short: "Create a NetBox device role",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateDeviceRole(arguments[0]))
		},
	}
}
