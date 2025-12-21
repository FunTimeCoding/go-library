package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
)

func GitTag() string {
	p := git.FindDirectory()
	result := git.LatestTag(p)

	if result == "" {
		fmt.Printf("No tag found: %s\n", p)
	}

	return result
}
