package convert

import "github.com/funtimecoding/go-library/pkg/prometheus/rule"

func Rules(v []*rule.Rule) []*SlimRule {
	var result []*SlimRule

	for _, r := range v {
		result = append(result, Rule(r))
	}

	return result
}
