package envelope

import "encoding/json"

type Envelope struct {
	Status       string          `json:"status"`
	Payload      json.RawMessage `json:"response"`
	Message      string          `json:"errorMessage"`
	InnerMessage string          `json:"innerErrorMessage"`
}
