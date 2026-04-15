package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func searchPagesCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "search-pages [query]",
		Short: "Search Confluence pages using CQL or plain text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.SearchPages(arguments[0]))
		},
	}
}
