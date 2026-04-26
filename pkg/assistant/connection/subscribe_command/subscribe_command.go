package subscribe_command

// SubscribeCommand is the payload for a WebSocket event subscription.
type SubscribeCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
	Event      string `json:"event_type,omitempty"`
}
