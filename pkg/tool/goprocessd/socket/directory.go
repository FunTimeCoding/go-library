package socket

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func directory() string {
	return filepath.Join(
		system.Home(),
		".local",
		"share",
		"goprocessd",
	)
}
