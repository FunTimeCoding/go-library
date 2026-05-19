package notation

import "encoding/json"

type ContentBlock struct {
	Type       string          `json:"type"`
	Identifier string          `json:"id"`
	Name       string          `json:"name"`
	Input      json.RawMessage `json:"input"`
}
