package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createTagCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tag [name]",
		Short: "Create a NetBox tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateTag(arguments[0]))
		},
	}
}
