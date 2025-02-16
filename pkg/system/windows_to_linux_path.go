package system

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func WindowsToLinuxPath(windowsPath string) string {
	result := strings.ReplaceAll(windowsPath, "\\", separator.Slash)

	if len(result) > 1 && result[1] == ':' {
		driveLetter := strings.ToLower(string(result[0]))
		result = separator.Slash + driveLetter + result[2:]
	}

	return result
}
