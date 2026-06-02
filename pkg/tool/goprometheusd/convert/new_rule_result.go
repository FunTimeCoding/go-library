package convert

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/paginate"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule"
)

func NewRuleResult(
	v []*rule.Rule,
	limit int,
	offset int,
) *SlimRuleResult {
	total := len(v)

	return &SlimRuleResult{
		Rules: paginate.Slice(Rules(v), limit, offset),
		Total: total,
	}
}
