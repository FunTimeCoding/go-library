package message

import "encoding/json"

type Event struct {
	Type    string          `json:"event_type"`
	Time    string          `json:"time_fired"`
	Origin  string          `json:"origin"`
	Context Context         `json:"context"`
	Raw     json.RawMessage `json:"data"`
}
