package constant

import "github.com/gorilla/websocket"

const (
	PluginEnvironment = "MONITOR_PLUGINS"

	GoAlert     = "goalert"
	GoContainer = "gocontainer"
	GoFile      = "gofile"
	GoGenie     = "gogenie"
	GoImage     = "goimage"
	GoJira      = "gojira"
	GoKevt      = "gokevt"
	GoSensor    = "gosensor"
	GoSentry    = "gosentry"
	GoSilence   = "gosilence"

	AlertPrefix           = "alert"
	ExamplePrefix         = "example"
	FilePrefix            = "file"
	GitHubPrefix          = "github"
	GitLabPrefix          = "gitlab"
	JiraPrefix            = "jira"
	KubernetesEventPrefix = "event"
	OpsgeniePrefix        = "opsgenie"
	PodManPrefix          = "podman"
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
