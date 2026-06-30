package vocabulary

import (
	"os"
	"path/filepath"
)

func cachePath() string {
	home, e := os.UserHomeDir()

	if e != nil {
		return ""
	}

	return filepath.Join(
		home,
		".local",
		"share",
		"tokenizer",
		"anthropic.json",
	)
}
