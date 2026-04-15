package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func transitionIssueCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "transition-issue [key] [transition-identifier]",
		Short: "Transition a Jira issue to a new status",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.TransitionIssue(arguments[0], arguments[1]))
		},
	}
}
