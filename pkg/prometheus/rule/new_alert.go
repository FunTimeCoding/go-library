package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

func NewAlert(
	r *v1.AlertingRule,
	g *v1.RuleGroup,
) *Rule {
	result := &Rule{
		Group:    g.Name,
		Name:     r.Name,
		RawAlert: r,
		RawGroup: g,
	}
	result.readAnnotations()

	return result
}
