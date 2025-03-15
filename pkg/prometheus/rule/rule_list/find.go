package rule_list

import "github.com/funtimecoding/go-library/pkg/prometheus/rule"

func (l *List) Find(name string) *rule.Rule {
	for _, r := range l.rules {
		if r.Name == name {
			return r
		}
	}

	return nil
}
