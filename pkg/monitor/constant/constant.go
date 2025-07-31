package constant

import "github.com/gorilla/websocket"

const (
	PluginEnvironment = "MONITOR_PLUGINS"
	FileEnvironment   = "MONITOR_FILE"
	ManualEnvironment = "MONITOR_MANUAL"

	Address = "localhost:8080"

	NotationReport int = 10 // Limit
)

// Command
const (
	LoginCommand  = "login"
	LogoutCommand = "logout"
	FlagCommand   = "flag"
	ClearCommand  = "clear"
	PingCommand   = "ping"

	LoginResponseCommand = "login-response"
	FlagAddCommand       = "flag-add"
	FlagRemoveCommand    = "flag-remove"
)

type (
	Severity string
	Status   string
)

const (
	Critical    Severity = "critical"
	Warning     Severity = "warning"
	Information Severity = "information"
)

const (
	Open       Status = "open"
	InProgress Status = "in-progress"
	Resolved   Status = "resolved"
	Closed     Status = "closed"
)

var (
	Upgrader = websocket.Upgrader{}

	Severities = []Severity{Critical, Warning, Information}

	Statuses = []Status{Open, InProgress, Resolved, Closed}
)
