package response

type EventEntryData struct {
	Values    []ExceptionValue `json:"values"`
	Message   string           `json:"message"`
	Formatted string           `json:"formatted"`
}
