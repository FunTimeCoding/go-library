package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func createSiteCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-site [name]",
		Short: "Create a NetBox site",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, arguments []string) {
			fmt.Println(c.CreateSite(arguments[0]))
		},
	}
}
