package issue

import "github.com/xanzy/go-gitlab"

type Issue struct {
	Project    int
	Identifier int
	Title      string
	State      string
	Link       string

	Raw *gitlab.Issue
}
