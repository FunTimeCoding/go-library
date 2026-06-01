package convert

type SlimAlert struct {
	Name        string            `json:"name"`
	State       string            `json:"state"`
	Severity    string            `json:"severity"`
	Summary     string            `json:"summary"`
	Fingerprint string            `json:"fingerprint"`
	Start       string            `json:"start"`
	Receivers   []string          `json:"receivers"`
	Labels      map[string]string `json:"labels,omitempty"`
	Link        string            `json:"link,omitempty"`
}
