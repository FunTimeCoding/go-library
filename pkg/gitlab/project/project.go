package project

import "github.com/xanzy/go-gitlab"

type Project struct {
	Identifier int
	Namespace  string
	Name       string

	Raw *gitlab.Project
}
