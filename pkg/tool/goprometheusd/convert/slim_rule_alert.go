package convert

type SlimRuleAlert struct {
	State       string            `json:"state"`
	ActiveAt    string            `json:"active_at"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Value       string            `json:"value"`
}
