package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func getIssueCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-issue [key]",
		Short: "Get a Jira issue by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetIssue(arguments[0]))
		},
	}
}
