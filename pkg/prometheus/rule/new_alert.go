package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

func NewAlert(
	u *v1.AlertingRule,
	g *v1.RuleGroup,
) *Rule {
	r := &Rule{
		Group:    g.Name,
		Name:     u.Name,
		Query:    u.Query,
		Health:   u.Health,
		State:    u.State,
		RawAlert: u,
		RawGroup: g,
	}

	return r.parse()
}
