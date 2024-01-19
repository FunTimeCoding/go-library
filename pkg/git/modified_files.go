package git

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/go-git/go-git/v5"
)

func ModifiedFiles() []string {
	var result []string

	for file, status := range Status(system.WorkingDirectory()) {
		// Staging: staged changes
		// Worktree: non-staged changes
		if status.Staging == git.Modified || status.Worktree == git.Modified {
			result = append(result, file)
		}
	}

	return result
}
