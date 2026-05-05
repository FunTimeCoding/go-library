package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createCluster(c *client.Client) *cobra.Command {
	var clusterType string
	var site string
	command := &cobra.Command{
		Use:   "create-cluster [name]",
		Short: "Create a NetBox cluster",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateCluster(arguments[0], clusterType, site))
		},
	}
	command.Flags().StringVar(
		&clusterType,
		"type",
		"",
		"cluster type name (required)",
	)
	command.Flags().StringVar(&site, "site", "", "site name (required)")
	errors.PanicOnError(command.MarkFlagRequired("type"))
	errors.PanicOnError(command.MarkFlagRequired("site"))

	return command
}
