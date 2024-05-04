package build

import "github.com/funtimecoding/go-library/pkg/git"

func GitTag() string {
	return git.LatestTag(git.FindDirectory())
}
