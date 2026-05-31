package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goclaude",
	"Claude Code session management and analysis",
	"goclaude [command]",
)

const (
	SessionIdentifierEnvironment = "CLAUDE_SESSION_ID"
	NameEnvironment              = "CLAUDE_NAME"
	HostEnvironment              = "CLAUDE_HOST"
	PortEnvironment              = "CLAUDE_PORT"
	PeekOutputBudget             = 120
)
