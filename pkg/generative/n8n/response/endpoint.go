package response

type Endpoint struct {
	Node  string `json:"node"`
	Type  string `json:"type"`
	Index int    `json:"index"`
}
