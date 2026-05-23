package usage_entry

func New(
	timestamp string,
	model string,
	inputTokens int,
	outputTokens int,
	cacheCreationInputTokens int,
	cacheReadInputTokens int,
) *Entry {
	return &Entry{
		Timestamp:                timestamp,
		Model:                    model,
		InputTokens:              inputTokens,
		OutputTokens:             outputTokens,
		CacheCreationInputTokens: cacheCreationInputTokens,
		CacheReadInputTokens:     cacheReadInputTokens,
	}
}
