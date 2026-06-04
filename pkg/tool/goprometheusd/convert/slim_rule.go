package convert

type SlimRule struct {
	Name           string          `json:"name"`
	Group          string          `json:"group"`
	Type           string          `json:"type"`
	State          string          `json:"state,omitempty"`
	Health         string          `json:"health"`
	Query          string          `json:"query"`
	Severity       string          `json:"severity,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	LastEvaluation string          `json:"last_evaluation,omitempty"`
	EvaluationTime float64         `json:"evaluation_time,omitempty"`
	LastError      string          `json:"last_error,omitempty"`
	Alerts         []SlimRuleAlert `json:"alerts,omitempty"`
}
