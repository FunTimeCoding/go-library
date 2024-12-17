package project

import "gitlab.com/gitlab-org/api/client-go"

type Project struct {
	Identifier int
	Namespace  string
	Name       string
	Link       string

	Raw *gitlab.Project
}
