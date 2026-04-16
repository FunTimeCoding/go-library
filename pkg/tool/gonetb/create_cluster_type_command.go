package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createClusterTypeCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cluster-type [name]",
		Short: "Create a NetBox cluster type",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateClusterType(arguments[0]))
		},
	}
}
