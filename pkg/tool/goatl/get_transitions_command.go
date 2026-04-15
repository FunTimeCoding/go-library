package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func getTransitionsCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-transitions [key]",
		Short: "List available transitions for a Jira issue",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetTransitions(arguments[0]))
		},
	}
}
