package goclaude

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
)

func formatCheckContext(r *client.CheckResponse) string {
	hasTimeout := r.TimeoutMessage != nil && *r.TimeoutMessage != ""
	hasMemory := r.MemoryActivity != nil && len(*r.MemoryActivity) > 0
	hasReannounce := r.Reannounce != nil && *r.Reannounce
	hasPulses := r.Pulses != nil && len(*r.Pulses) > 0

	if len(r.Sessions) == 0 &&
		len(r.Messages) == 0 &&
		len(r.Completions) == 0 &&
		!hasTimeout &&
		!hasMemory &&
		!hasReannounce &&
		!hasPulses {
		return ""
	}

	parts := []string{
		fmt.Sprintf("[goclauded] Called %s today.", r.Callsign),
	}

	if hasReannounce {
		parts = append(
			parts,
			"Re-announce required: MCP binding lost during service restart. Call announce with your session name and topic to restore.",
		)
	}

	if hasPulses {
		for _, p := range *r.Pulses {
			parts = append(
				parts,
				fmt.Sprintf("[pulse] %s", p.Body),
			)
		}
	}

	if len(r.Sessions) > 0 {
		parts = append(parts, "Active sessions:")

		for _, s := range r.Sessions {
			line := fmt.Sprintf("  %s", s.Callsign)

			if s.Topic != "" {
				line = fmt.Sprintf("%s - %s", line, s.Topic)
			}

			parts = append(parts, line)

			if s.Labels != nil && len(*s.Labels) > 0 {
				var pips []string

				for _, l := range *s.Labels {
					pips = append(
						pips,
						fmt.Sprintf("(%s:%s)", l.Key, l.Value),
					)
				}

				parts = append(
					parts,
					fmt.Sprintf("    %s", join.Space(pips...)),
				)
			}
		}
	}

	if len(r.Completions) > 0 {
		parts = append(parts, "Recent activity:")

		for _, c := range r.Completions {
			switch c.Kind {
			case "complete":
				parts = append(
					parts,
					fmt.Sprintf("  %s completed %s", c.Name, c.Topic),
				)
			case "update":
				parts = append(
					parts,
					fmt.Sprintf(
						"  %s updated scope: %s",
						c.Name,
						c.Topic,
					),
				)
			}
		}
	}

	if hasTimeout {
		parts = append(
			parts,
			fmt.Sprintf("Timeout: %s", *r.TimeoutMessage),
		)
	}

	if hasMemory {
		parts = append(parts, "Memory activity:")

		for _, a := range *r.MemoryActivity {
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

	if len(r.Messages) > 0 {
		parts = append(parts, "Messages:")

		for _, m := range r.Messages {
			parts = append(
				parts,
				fmt.Sprintf("  %s: %s", m.From, m.Body),
			)
		}
	}

	return join.NewLine(parts)
}
