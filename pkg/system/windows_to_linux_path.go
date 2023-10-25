package system

import "strings"

func WindowsToLinuxPath(windowsPath string) string {
	result := strings.ReplaceAll(windowsPath, "\\", "/")

	if len(result) > 1 && result[1] == ':' {
		driveLetter := strings.ToLower(string(result[0]))
		result = "/" + driveLetter + result[2:]
	}

	return result
}
