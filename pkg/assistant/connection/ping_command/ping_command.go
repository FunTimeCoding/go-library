package ping_command

// PingCommand is the payload for a WebSocket ping.
type PingCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
}
