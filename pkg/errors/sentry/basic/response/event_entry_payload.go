package response

import "encoding/json"

type EventEntryPayload struct {
	Values      []ExceptionValue  `json:"values"`
	Message     string            `json:"message"`
	Formatted   string            `json:"formatted"`
	Locator     string            `json:"url"`
	Method      string            `json:"method"`
	Payload     json.RawMessage   `json:"data"`
	Headers     [][]string        `json:"headers"`
	Query       string            `json:"query"`
	Cookies     [][]string        `json:"cookies"`
	Environment map[string]string `json:"env"`
}
