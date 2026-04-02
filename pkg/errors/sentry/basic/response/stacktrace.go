package response

type Stacktrace struct {
	Frames []StackFrame `json:"frames"`
}
