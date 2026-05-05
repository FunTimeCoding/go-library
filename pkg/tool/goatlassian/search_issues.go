package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func searchIssues(c *client.Client) *cobra.Command {
	var limit int
	command := &cobra.Command{
		Use:   "search-issues [query]",
		Short: "Search Jira issues using JQL",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var l *int

			if limit > 0 {
				l = &limit
			}

			fmt.Println(c.SearchIssues(arguments[0], l))
		},
	}
	command.Flags().IntVar(&limit, "limit", 0, "maximum number of results")

	return command
}
