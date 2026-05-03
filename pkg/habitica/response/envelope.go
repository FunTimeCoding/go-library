package response

import "encoding/json"

type Envelope struct {
	Payload json.RawMessage `json:"data"`
}
