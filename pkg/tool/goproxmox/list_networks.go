package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
)

func listNetworks(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-networks [node]",
		Short: "List network interfaces on a node",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			fmt.Println(c.ListNetworks(a[0]))
		},
	}
}
