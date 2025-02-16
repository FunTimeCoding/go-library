package constant

import "github.com/gorilla/websocket"

const (
	GoSensor = "gosensor"
	GoSentry = "gosentry"

	ExamplePrefix = "example"
	SentryPrefix  = "sentry"

	ErrorType = "error"

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
