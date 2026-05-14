package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
)

func listNodes(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-nodes",
		Short: "List Proxmox nodes",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListNodes())
		},
	}
}
