package ssh

type Result struct {
	OutputString string
	ErrorString  string
	ExitCode     int
	Error        error
}
