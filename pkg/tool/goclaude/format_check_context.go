package goclaude

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"strings"
)

func formatCheckContext(body *client.CheckResponse) string {
	hasTimeout := body.TimeoutMessage != nil && *body.TimeoutMessage != ""
	hasMemory := body.MemoryActivity != nil && len(*body.MemoryActivity) > 0
	hasReannounce := body.Reannounce != nil && *body.Reannounce

	if len(body.Sessions) == 0 && len(body.Messages) == 0 && len(body.Completions) == 0 && !hasTimeout && !hasMemory && !hasReannounce {
		return ""
	}

	var parts []string
	parts = append(
		parts,
		fmt.Sprintf("[goclauded] Called %s today.", body.Name),
	)

	if hasReannounce {
		parts = append(
			parts,
			"Re-announce required: MCP binding lost during service restart. Call announce with your session name and topic to restore.",
		)
	}

	if len(body.Sessions) > 0 {
		parts = append(parts, "Active sessions:")

		for _, s := range body.Sessions {
			line := fmt.Sprintf("  %s", s.Name)

			if s.Topic != "" {
				line = fmt.Sprintf("%s - %s", line, s.Topic)
			}

			parts = append(parts, line)
		}
	}

	if len(body.Completions) > 0 {
		parts = append(parts, "Recent activity:")

		for _, c := range body.Completions {
			switch c.Kind {
			case "complete":
				parts = append(
					parts,
					fmt.Sprintf("  %s completed %s", c.Name, c.Topic),
				)
			case "update":
				parts = append(
					parts,
					fmt.Sprintf("  %s updated scope: %s", c.Name, c.Topic),
				)
			}
		}
	}

	if hasTimeout {
		parts = append(
			parts,
			fmt.Sprintf("Timeout: %s", *body.TimeoutMessage),
		)
	}

	if hasMemory {
		parts = append(parts, "Memory activity:")

		for _, a := range *body.MemoryActivity {
			if a.Source != "" {
				parts = append(
					parts,
					fmt.Sprintf(
						"  %s %s by %s",
						a.Name,
						a.ChangeType,
						a.Source,
					),
				)
			} else {
				parts = append(
					parts,
					fmt.Sprintf("  %s %s", a.Name, a.ChangeType),
				)
			}
		}
	}

	if len(body.Messages) > 0 {
		parts = append(parts, "Messages:")

		for _, m := range body.Messages {
			parts = append(
				parts,
				fmt.Sprintf(
					"  %s: %s",
					m.From,
					m.Body,
				),
			)
		}
	}

	return strings.Join(parts, "\n")
}
