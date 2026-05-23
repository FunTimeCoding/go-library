package usage

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/usage_entry"

func entryCost(
	model string,
	e *usage_entry.Entry,
) float64 {
	var inputRate float64
	var outputRate float64
	var cacheCreateRate float64
	var cacheReadRate float64

	switch model {
	case "opus":
		inputRate = 15.0
		outputRate = 75.0
		cacheCreateRate = 18.75
		cacheReadRate = 1.5
	case "sonnet":
		inputRate = 3.0
		outputRate = 15.0
		cacheCreateRate = 3.75
		cacheReadRate = 0.3
	case "haiku":
		inputRate = 0.25
		outputRate = 1.25
		cacheCreateRate = 0.3
		cacheReadRate = 0.03
	default:
		inputRate = 3.0
		outputRate = 15.0
		cacheCreateRate = 3.75
		cacheReadRate = 0.3
	}

	return float64(e.InputTokens)/1_000_000*inputRate +
		float64(e.OutputTokens)/1_000_000*outputRate +
		float64(e.CacheCreationInputTokens)/1_000_000*
			cacheCreateRate +
		float64(e.CacheReadInputTokens)/1_000_000*
			cacheReadRate
}
