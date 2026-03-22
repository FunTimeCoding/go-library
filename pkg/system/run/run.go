package run

type Run struct {
	environment  []string
	stdio        bool
	Directory    string
	Panic        bool
	Verbose      bool
	OutputString string
	ErrorString  string
	Error        error
	Exit         int
}
