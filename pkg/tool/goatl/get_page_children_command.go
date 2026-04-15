package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func getPageChildrenCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-page-children [identifier]",
		Short: "List child pages of a Confluence page",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetPageChildren(arguments[0]))
		},
	}
}
