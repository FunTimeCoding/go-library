package constant

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func SocketDirectory() string {
	return filepath.Join(
		system.Home(),
		".local",
		"share",
		"goprocessd",
	)
}
