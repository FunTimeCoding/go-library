package run

type Run struct {
	Environment []string
	Directory   string

	Panic   bool
	Verbose bool

	OutputString string
	ErrorString  string
	Error        error
}
