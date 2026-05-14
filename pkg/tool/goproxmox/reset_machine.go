package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/client"
	"github.com/spf13/cobra"
	"strconv"
)

func resetMachine(c *client.Client) *cobra.Command {
	var node string
	result := &cobra.Command{
		Use:   "reset-machine [vmid]",
		Short: "Reset a virtual machine",
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

			fmt.Println(c.ResetMachine(vmid, n))
		},
	}
	result.Flags().StringVar(
		&node,
		"node",
		"",
		"node name (speeds up lookup)",
	)

	return result
}
