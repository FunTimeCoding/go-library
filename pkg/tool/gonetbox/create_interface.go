package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createInterface(c *client.Client) *cobra.Command {
	var interfaceType string
	command := &cobra.Command{
		Use:   "create-interface [device] [name]",
		Short: "Create a network interface on a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateInterface(
					arguments[0],
					arguments[1],
					interfaceType,
				),
			)
		},
	}
	command.Flags().StringVar(
		&interfaceType,
		"type",
		"",
		"interface type (e.g. 1000base-t)",
	)
	errors.PanicOnError(command.MarkFlagRequired("type"))

	return command
}
