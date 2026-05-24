package goclaude

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/writer"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
	"os"
)

func register(c *command_context.Context) *cobra.Command {
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

			response, e := c.Client().PostRegisterWithResponse(
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

			name := response.JSON200.Callsign
			envFile := os.Getenv("CLAUDE_ENV_FILE")

			if envFile != "" {
				f, fe := os.OpenFile(
					envFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY,
					0644,
				)

				if fe == nil {
					writer.Print(
						f,
						"export %s=%s\n",
						constant.SessionIdentifierEnvironment,
						sessionID,
					)
					writer.Print(
						f,
						"export %s=%s\n",
						constant.NameEnvironment,
						name,
					)
					errors.PanicOnError(f.Close())
				}
			}
		},
	}

	return result
}
