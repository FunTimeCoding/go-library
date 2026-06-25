package record

type Record struct {
	Disabled bool              `json:"disabled"`
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	TTL      int               `json:"ttl"`
	Payload  map[string]any    `json:"rData"`
	Comments string            `json:"comments"`
}
