package build

import "fmt"

func (b *Build) Print() {
	fmt.Printf(
		"Version: %s\nGitHash: %s\nBuildDate: %s\n",
		b.Version,
		b.GitHash,
		b.BuildDate,
	)
}
