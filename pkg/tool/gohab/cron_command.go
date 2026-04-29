package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
	"github.com/spf13/cobra"
)

func cronCommand(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "cron",
		Short: "Check and run daily rollover if needed",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Cron())
		},
	}
}
