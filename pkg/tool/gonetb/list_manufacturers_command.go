package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/spf13/cobra"
)

func listManufacturersCommand(c *netb.Client) *cobra.Command {
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
