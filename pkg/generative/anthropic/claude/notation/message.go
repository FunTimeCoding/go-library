package notation

import "encoding/json"

type Message struct {
	Role    string          `json:"role"`
	Content json.RawMessage `json:"content"`
	Model   string          `json:"model"`
	Usage   *MessageUsage   `json:"usage"`
}
