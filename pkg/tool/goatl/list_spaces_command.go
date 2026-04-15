package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func listSpacesCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-spaces",
		Short: "List all visible Confluence spaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListSpaces())
		},
	}
}
