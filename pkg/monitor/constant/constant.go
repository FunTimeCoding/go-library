package constant

import "github.com/gorilla/websocket"

const (
	PluginEnvironment = "MONITOR_PLUGINS"
	FileEnvironment   = "MONITOR_FILE"
	ManualEnvironment = "MONITOR_MANUAL"

	GoAlert     = "goalert"
	GoContainer = "gocontainer"
	GoFile      = "gofile"
	GoGenie     = "gogenie"
	GoGitHub    = "gogithub"
	GoGitLab    = "gogitlab"
	GoGitStatus = "gogitstatus"
	GoImage     = "goimage"
	GoJira      = "gojira"
	GoKevt      = "gokevt"
	GoSentry    = "gosentry"
	GoSilence   = "gosilence"
	GoVersion   = "goversion"
	GoVolume    = "govolume"

	AlertPrefix           = "alert"
	ExamplePrefix         = "example"
	FilePrefix            = "file"
	GitHubPrefix          = "github"
	GitLabPrefix          = "gitlab"
	GitPrefix             = "git"
	GoPrefix              = "go"
	JiraPrefix            = "jira"
	KubernetesEventPrefix = "event"
	MonitorPrefix         = "monitor"
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
