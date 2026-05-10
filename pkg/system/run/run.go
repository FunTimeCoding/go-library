package run

import "github.com/funtimecoding/go-library/pkg/face"

type Run struct {
	environment   []string
	stdio         bool
	reporter      face.Reporter
	reporterLabel string
	Directory     string
	Panic         bool
	Verbose       bool
	OutputString  string
	ErrorString   string
	Error         error
	Exit          int
}
