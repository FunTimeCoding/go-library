package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/spf13/cobra"
)

func tagsCommand(c *habitica.Client) *cobra.Command {
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
