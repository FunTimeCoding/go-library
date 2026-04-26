package response

type EventEntryPayload struct {
	Values    []ExceptionValue `json:"values"`
	Message   string           `json:"message"`
	Formatted string           `json:"formatted"`
}
