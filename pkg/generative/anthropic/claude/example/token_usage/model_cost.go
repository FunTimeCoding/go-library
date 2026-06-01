package token_usage

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/example/common"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/usage_entry"
)

func modelCost(model string, m *modelTotals) float64 {
	return common.EntryCost(
		model,
		usage_entry.New(
			"",
			model,
			m.input,
			m.output,
			m.cacheCreate,
			m.cacheRead,
		),
	)
}
