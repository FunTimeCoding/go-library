package usage

import "strings"

func normalizeModel(model string) string {
	lower := strings.ToLower(model)

	if strings.Contains(lower, "opus") {
		return "opus"
	}

	if strings.Contains(lower, "sonnet") {
		return "sonnet"
	}

	if strings.Contains(lower, "haiku") {
		return "haiku"
	}

	return model
}
