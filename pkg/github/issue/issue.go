package issue

import "github.com/google/go-github/v82/github"

type Issue struct {
	Repository string
	Title      string
	Link       string
	Raw        *github.Issue
}
