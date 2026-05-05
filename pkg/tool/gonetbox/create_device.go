package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDevice(c *client.Client) *cobra.Command {
	var role string
	var deviceType string
	var site string
	var tenant string
	command := &cobra.Command{
		Use:   "create-device [name]",
		Short: "Create a NetBox device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var t *string

			if tenant != "" {
				t = &tenant
			}

			fmt.Println(
				c.CreateDevice(
					arguments[0],
					role,
					deviceType,
					site,
					t,
				),
			)
		},
	}
	command.Flags().StringVar(
		&role,
		"role",
		"",
		"device role name (required)",
	)
	command.Flags().StringVar(
		&deviceType,
		"type",
		"",
		"device type model (required)",
	)
	command.Flags().StringVar(&site, "site", "", "site name (required)")
	command.Flags().StringVar(&tenant, "tenant", "", "tenant name (optional)")
	errors.PanicOnError(command.MarkFlagRequired("role"))
	errors.PanicOnError(command.MarkFlagRequired("type"))
	errors.PanicOnError(command.MarkFlagRequired("site"))

	return command
}
