package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
	"strconv"
)

func createSnapshot(c *client.Client) *cobra.Command {
	var node string
	var name string
	result := &cobra.Command{
		Use:   "create-snapshot [vmid]",
		Short: "Create a snapshot of a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			vmid, e := strconv.ParseInt(a[0], 10, 64)
			errors.PanicOnError(e)
			var n *string

			if node != "" {
				n = &node
			}

			fmt.Println(c.CreateSnapshot(vmid, name, n))
		},
	}
	result.Flags().StringVar(
		&name,
		"name",
		"",
		"snapshot name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("name"))
	result.Flags().StringVar(
		&node,
		"node",
		"",
		"node name (speeds up lookup)",
	)

	return result
}
