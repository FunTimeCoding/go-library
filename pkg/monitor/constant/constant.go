package constant

import "github.com/gorilla/websocket"

const (
	PluginEnvironment = "MONITOR_PLUGINS"

	GoAlert   = "goalert"
	GoFile    = "gofile"
	GoGenie   = "gogenie"
	GoJira    = "gojira"
	GoKevt    = "gokevt"
	GoSensor  = "gosensor"
	GoSentry  = "gosentry"
	GoSilence = "gosilence"

	AlertPrefix           = "alert"
	ExamplePrefix         = "example"
	FilePrefix            = "file"
	JiraPrefix            = "jira"
	KubernetesEventPrefix = "event"
	OpsgeniePrefix        = "opsgenie"
	SentryPrefix          = "sentry"
	SilencePrefix         = "silence"

	InformationLevel = "information"
	WarningLevel     = "warning"
	ErrorLevel       = "error"

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

const NotationReport int = 10
