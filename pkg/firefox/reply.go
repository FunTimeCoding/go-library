package firefox

import "encoding/json"

type reply struct {
	Result     json.RawMessage `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	Identifier int             `json:"id"`
}
