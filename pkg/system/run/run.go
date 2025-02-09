package run

type Run struct {
	environment []string
	Directory   string

	Panic   bool
	Verbose bool

	OutputString string
	ErrorString  string
	Error        error
	ExitCode     int
}
