package issue

import "gitlab.com/gitlab-org/api/client-go"

type Issue struct {
	Project    int
	Identifier int
	Title      string
	State      string
	Link       string

	Raw *gitlab.Issue
}
