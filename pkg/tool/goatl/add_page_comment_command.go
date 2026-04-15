package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/spf13/cobra"
)

func addPageCommentCommand(c *atl.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-page-comment [identifier] [body]",
		Short: "Add a comment to a Confluence page",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.AddPageComment(arguments[0], arguments[1]))
		},
	}
}
