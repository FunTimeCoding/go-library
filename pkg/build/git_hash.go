package build

import "github.com/funtimecoding/go-library/pkg/git"

func GitHash() string {
	return git.ShortHash(git.FindDirectory())
}
