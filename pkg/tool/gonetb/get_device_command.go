package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func getDeviceCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-device [name]",
		Short: "Get a NetBox device by name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetDevice(arguments[0]))
		},
	}
}
