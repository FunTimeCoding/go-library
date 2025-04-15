package rule_list

import "github.com/prometheus/client_golang/api/prometheus/v1"

func LabelsSame(
	a *v1.AlertingRule,
	b *v1.AlertingRule,
) bool {
	if len(a.Labels) != len(b.Labels) {
		return false
	}

	for k, v := range a.Labels {
		if b.Labels[k] != v {
			return false
		}
	}

	return true
}
