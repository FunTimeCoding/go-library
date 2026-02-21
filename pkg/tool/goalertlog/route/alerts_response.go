package route

type AlertsResponse struct {
	Fingerprint string            `json:"fingerprint"`
	Name        string            `json:"name"`
	Severity    string            `json:"severity"`
	Summary     string            `json:"summary"`
	Labels      map[string]string `json:"labels"`
	Start       string            `json:"start"`
	End         string            `json:"end,omitempty"`
	Status      string            `json:"status"`
}
