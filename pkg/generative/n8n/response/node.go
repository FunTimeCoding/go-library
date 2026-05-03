package response

type Node struct {
	Parameters  Parameters     `json:"parameters"`
	Type        string         `json:"type"`
	TypeVersion float64        `json:"typeVersion"`
	Position    []int          `json:"position"`
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Credentials map[string]any `json:"credentials,omitempty"`
	WebhookId   string         `json:"webhookId,omitempty"`
}
