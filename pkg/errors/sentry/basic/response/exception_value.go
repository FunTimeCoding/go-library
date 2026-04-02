package response

type ExceptionValue struct {
	Type       string      `json:"type"`
	Value      string      `json:"value"`
	Stacktrace *Stacktrace `json:"stacktrace"`
}
