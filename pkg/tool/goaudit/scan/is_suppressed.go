package scan

func isSuppressed(
	suppress map[string][]string,
	operationID string,
	concernKey string,
) bool {
	keys, okay := suppress[operationID]

	if !okay {
		return false
	}

	for _, k := range keys {
		if k == concernKey {
			return true
		}
	}

	return false
}
