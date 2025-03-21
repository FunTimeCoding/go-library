package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

func NewRecord(
	r *v1.RecordingRule,
	g *v1.RuleGroup,
) *Rule {
	result := &Rule{
		Group:     g.Name,
		Name:      r.Name,
		RawRecord: r,
		RawGroup:  g,
	}

	return result.parse()
}
