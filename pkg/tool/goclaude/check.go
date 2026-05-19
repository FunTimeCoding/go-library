package goclaude

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func check() *cobra.Command {
	var host string
	var port int
	result := &cobra.Command{
		Use:   "check",
		Short: "Check goclauded for session roster and messages (hook output)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			sessionID := sessionIdentifierFromEnvironment()
			c, e := client.NewClientWithResponses(
				locator.New(host).Port(port).Insecure().String(),
			)
			errors.PanicOnError(e)
			response, e := c.GetCheckWithResponse(
				context.Background(),
				&client.GetCheckParams{Session: sessionID},
			)

			if e != nil {
				os.Exit(1)
			}

			if response.JSON200 == nil {
				os.Exit(1)
			}

			body := response.JSON200

			if !body.Changed {
				return
			}

			context := formatCheckContext(body)

			if context == "" {
				return
			}

			output := hookOutput{
				HookSpecificOutput: hookSpecificOutput{
					HookEventName:     "UserPromptSubmit",
					AdditionalContext: context,
				},
			}
			errors.PanicOnError(json.NewEncoder(os.Stdout).Encode(output))
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
