package result

type Result struct {
	OutputString string
	ErrorString  string
	ExitCode     int
	Error        error
}
