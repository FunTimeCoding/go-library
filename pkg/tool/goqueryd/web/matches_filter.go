package web

func matchesFilter(
	documentMetadata map[string]string,
	filter map[string]string,
) bool {
	for key, value := range filter {
		if documentMetadata == nil || documentMetadata[key] != value {
			return false
		}
	}

	return true
}
