package message

import "encoding/json"

type Message struct {
	Identifier uint64          `json:"id"`
	Type       string          `json:"type"`
	Success    bool            `json:"success"`
	Result     json.RawMessage `json:"result"`
	Event      *Event          `json:"event,omitempty"`
	Error      *Error          `json:"error,omitempty"`
}
