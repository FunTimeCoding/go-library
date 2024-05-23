package command

type Command struct {
	Command         string
	Environment     map[string]string
	RequestTeletype bool
}
