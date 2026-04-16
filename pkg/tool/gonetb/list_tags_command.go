package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listTagsCommand(c *netb.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all NetBox tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListTags())
		},
	}
}
