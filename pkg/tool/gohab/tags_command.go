package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
	"github.com/spf13/cobra"
)

func tagsCommand(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "tags",
		Short: "List tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Tags())
		},
	}
}
