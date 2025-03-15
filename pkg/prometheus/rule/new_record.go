package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

func NewRecord(
	r *v1.RecordingRule,
	g *v1.RuleGroup,
) *Rule {
	return &Rule{
		Group:     g.Name,
		Name:      r.Name,
		RawRecord: r,
		RawGroup:  g,
	}
}
