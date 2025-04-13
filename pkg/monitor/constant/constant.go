package constant

import "github.com/gorilla/websocket"

const (
	GoAlert   = "goalert"
	GoGenie   = "gogenie"
	GoJira    = "gojira"
	GoKevt    = "gokevt"
	GoSensor  = "gosensor"
	GoSentry  = "gosentry"
	GoSilence = "gosilence"

	AlertPrefix    = "alert"
	ExamplePrefix  = "example"
	JiraPrefix     = "jira"
	OpsgeniePrefix = "opsgenie"
	SentryPrefix   = "sentry"
	SilencePrefix  = "silence"

	WarningType = "warning"
	ErrorType   = "error"

	Address = "localhost:8080"

	LoginCommand  = "login"
	LogoutCommand = "logout"
	FlagCommand   = "flag"
	ClearCommand  = "clear"
	PingCommand   = "ping"

	LoginResponseCommand = "login-response"
	FlagAddCommand       = "flag-add"
	FlagRemoveCommand    = "flag-remove"
)

var Upgrader = websocket.Upgrader{}
