package issue

import "github.com/google/go-github/v69/github"

type Issue struct {
	Repository string
	Title      string
	Link       string
	Raw        *github.Issue
}
