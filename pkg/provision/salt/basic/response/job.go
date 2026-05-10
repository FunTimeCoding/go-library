package response

type Job struct {
	JID        string               `json:"jid,omitempty"`
	Function   string               `json:"Function,omitempty"`
	Arguments  []any                `json:"Arguments,omitempty"`
	Target     any                  `json:"Target,omitempty"`
	TargetType string               `json:"Target-Type,omitempty"`
	StartTime  string               `json:"StartTime"`
	User       string               `json:"User"`
	Minions    []string             `json:"Minions,omitempty"`
	Result     map[string]JobResult `json:"Result,omitempty"`
	Error      string               `json:"Error,omitempty"`
}
