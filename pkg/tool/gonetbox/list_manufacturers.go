package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listManufacturers(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-manufacturers",
		Short: "List all NetBox manufacturers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListManufacturers())
		},
	}
}
