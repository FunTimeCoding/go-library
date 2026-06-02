package argument

type ListAlerts struct {
	Filter      string  `json:"filter"`
	Active      string  `json:"active"`
	Silenced    string  `json:"silenced"`
	Inhibited   string  `json:"inhibited"`
	Unprocessed string  `json:"unprocessed"`
	Receiver    string  `json:"receiver"`
	Limit       float64 `json:"limit"`
	Offset      float64 `json:"offset"`
}
