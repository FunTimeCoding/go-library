package console

import "github.com/fatih/color"

var (
	CyanInstance    = color.New(color.FgCyan)
	GreenInstance   = color.New(color.FgGreen)
	MagentaInstance = color.New(color.FgMagenta)
	RedInstance     = color.New(color.FgRed)
	YellowInstance  = color.New(color.FgYellow)

	Cyan    = CyanInstance.SprintfFunc()
	Green   = GreenInstance.SprintfFunc()
	Magenta = MagentaInstance.SprintfFunc()
	Red     = RedInstance.SprintfFunc()
	Yellow  = YellowInstance.SprintfFunc()
)

const (
	GreenColor  = "green"
	RedColor    = "red"
	YellowColor = "yellow"
)
