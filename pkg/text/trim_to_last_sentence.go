package text

import "strings"

func TrimToLastSentence(text string) string {
	i := strings.LastIndexAny(text, ".!?")

	// No punctuation found, or already ends with punctuation
	if i == -1 || i == len(text)-1 {
		return text
	}

	// Trim to include the punctuation (changed i > 0 to i >= 0)
	if i >= 0 && i+1 < len(text) {
		return text[:i+1]
	}

	return text
}
