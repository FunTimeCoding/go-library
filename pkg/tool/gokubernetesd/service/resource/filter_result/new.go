package filter_result

func New(
	object map[string]interface{},
	filtered []string,
) *FilterResult {
	return &FilterResult{
		Object:   object,
		Filtered: filtered,
	}
}
