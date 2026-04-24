package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/spf13/cobra"
)

func statsCommand(c *habitica.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Get user stats",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Stats())
		},
	}
}
