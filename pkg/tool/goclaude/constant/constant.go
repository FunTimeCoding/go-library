package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goclaude",
	"Claude Code session management and analysis",
	"goclaude [command]",
)

const SessionIdentifierEnvironment = "GOCLAUDED_SESSION_ID"
