package build

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git"
)

func GitTag() string {
	p := git.FindDirectory()
	result := git.LatestTag(p)

	if result == "" {
		errors.Warning("no tag found: %s", p)
	}

	return result
}
