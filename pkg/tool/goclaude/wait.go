package goclaude

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func wait() *cobra.Command {
	var host string
	var port int
	var timeout int
	result := &cobra.Command{
		Use:   "wait",
		Short: "Long-poll goclauded for messages (asyncRewake hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			name := os.Getenv("GOCLAUDED_NAME")
			c, e := client.NewClientWithResponses(
				locator.New(host).Port(port).Insecure().String(),
			)
			errors.PanicOnError(e)
			response, e := c.GetWaitWithResponse(
				context.Background(),
				&client.GetWaitParams{
					Name:    name,
					Timeout: &timeout,
				},
			)

			if e != nil {
				os.Exit(0)
			}

			if response.JSON200 == nil {
				os.Exit(0)
			}

			messages := response.JSON200.Messages

			if len(messages) == 0 {
				os.Exit(0)
			}

			for _, m := range messages {
				errors.Printf("%s: %s\n", m.From, m.Body)
			}

			os.Exit(2)
		},
	}
	result.Flags().StringVar(
		&host,
		"host",
		"localhost",
		"goclauded host",
	)
	result.Flags().IntVar(
		&port,
		"port",
		8583,
		"goclauded port",
	)
	result.Flags().IntVar(
		&timeout,
		"timeout",
		30,
		"poll timeout in seconds",
	)

	return result
}
