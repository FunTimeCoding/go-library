package runner

import "github.com/xanzy/go-gitlab"

type Runner struct {
	Identifier  int
	Name        string
	Description string
	Address     string
	Status      string
	Type        string

	Raw *gitlab.Runner
}
