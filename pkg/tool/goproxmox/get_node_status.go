package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
)

func getNodeStatus(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-node-status [name]",
		Short: "Get Proxmox node status",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			fmt.Println(c.GetNodeStatus(a[0]))
		},
	}
}
