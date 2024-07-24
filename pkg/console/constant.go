package console

import "github.com/fatih/color"

var (
	Yellow  = color.New(color.FgYellow).SprintfFunc()
	Green   = color.New(color.FgGreen).SprintfFunc()
	Red     = color.New(color.FgRed).SprintfFunc()
	Cyan    = color.New(color.FgCyan).SprintfFunc()
	Magenta = color.New(color.FgMagenta).SprintfFunc()
)
