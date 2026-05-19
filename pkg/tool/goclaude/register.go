package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func register() *cobra.Command {
	var host string
	var port int
	result := &cobra.Command{
		Use:   "register",
		Short: "Register session with goclauded (SessionStart hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			input := readHookInput()
			sessionID := input.SessionID

			if sessionID == "" {
				return
			}

			c, e := client.NewClientWithResponses(
				locator.New(host).Port(port).Insecure().String(),
			)
			errors.PanicOnError(e)
			response, e := c.PostRegisterWithResponse(
				context.Background(),
				client.PostRegisterJSONRequestBody{
					Session: sessionID,
				},
			)

			if e != nil {
				return
			}

			if response.JSON200 == nil {
				return
			}

			name := response.JSON200.Name
			envFile := os.Getenv("CLAUDE_ENV_FILE")

			if envFile != "" {
				f, fe := os.OpenFile(
					envFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY,
					0644,
				)

				if fe == nil {
					_, e = fmt.Fprintf(
						f,
						"export GOCLAUDED_SESSION_ID=%s\n",
						sessionID,
					)
					errors.PanicOnError(e)
					_, e = fmt.Fprintf(f, "export GOCLAUDED_NAME=%s\n", name)
					errors.PanicOnError(e)
					errors.PanicOnError(f.Close())
				}
			}
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

	return result
}
