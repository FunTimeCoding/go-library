package git

import "github.com/funtimecoding/go-library/pkg/system"

func IsCleanCommand() bool {
	return system.Run("git", "status", "--porcelain") == ""
}
