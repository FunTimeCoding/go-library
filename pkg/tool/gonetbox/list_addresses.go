package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listAddresses(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-addresses [device]",
		Short: "List IP addresses for a NetBox device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.ListAddresses(arguments[0]))
		},
	}
}
