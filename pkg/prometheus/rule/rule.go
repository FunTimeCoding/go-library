package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

type Rule struct {
	Group         string
	Name          string
	Documentation string
	RawGroup      *v1.RuleGroup
	RawAlert      *v1.AlertingRule
	RawRecord     *v1.RecordingRule
}
