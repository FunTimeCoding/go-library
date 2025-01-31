package result

func New(
	outputString string,
	errorString string,
	code int,
	e error,
) *Result {
	return &Result{
		OutputString: outputString,
		ErrorString:  errorString,
		ExitCode:     code,
		Error:        e,
	}
}
