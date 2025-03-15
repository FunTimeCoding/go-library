package rule_list

import "github.com/funtimecoding/go-library/pkg/prometheus/rule"

func (l *List) Get() []*rule.Rule {
	return l.rules
}
