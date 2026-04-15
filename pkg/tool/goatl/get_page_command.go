package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func getPageCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-page [identifier]",
		Short: "Get a Confluence page by ID with body content",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetPage(arguments[0]))
		},
	}
}
