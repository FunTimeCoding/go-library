package sweep

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func sourcePath() string {
	return filepath.Join(
		system.Home(),
		".claude",
		"projects",
	)
}
