package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/cache"
)

func buildSystem(
	system string,
	m cache.Mode,
) []anthropic.TextBlockParam {
	if m == cache.None {
		return []anthropic.TextBlockParam{{Text: system}}
	}

	return []anthropic.TextBlockParam{
		{Text: system, CacheControl: cache.ToParameter(m)},
	}
}
