package container

import "github.com/google/go-github/v85/github"

type Container struct {
	Name       string
	Repository string
	Raw        *github.Package
}
