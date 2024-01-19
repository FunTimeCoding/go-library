package git

import "github.com/go-git/go-git/v5"

func ModifiedFiles(path string) []string {
	var result []string

	for file, status := range Status(path) {
		// Staging: staged changes
		// Worktree: non-staged changes
		if status.Staging == git.Modified || status.Worktree == git.Modified {
			result = append(result, file)
		}
	}

	return result
}
