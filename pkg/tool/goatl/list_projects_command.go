package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func listProjectsCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-projects",
		Short: "List all visible Jira projects",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListProjects())
		},
	}
}
