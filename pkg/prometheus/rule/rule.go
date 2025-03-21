package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

type Rule struct {
	Group       string
	Name        string
	Summary     string
	Description string
	Query       string
	Duration    int
	Runbook     string
	Severity    string
	State       string
	Health      v1.RuleHealth

	RawGroup  *v1.RuleGroup
	RawAlert  *v1.AlertingRule
	RawRecord *v1.RecordingRule
}
