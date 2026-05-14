package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
)

func listContainers(c *client.Client) *cobra.Command {
	var node string
	result := &cobra.Command{
		Use:   "list-containers",
		Short: "List LXC containers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var n *string

			if node != "" {
				n = &node
			}

			fmt.Println(c.ListContainers(n))
		},
	}
	result.Flags().StringVar(
		&node,
		"node",
		"",
		"filter by node name",
	)

	return result
}
