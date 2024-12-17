package runner

import "gitlab.com/gitlab-org/api/client-go"

type Runner struct {
	Identifier  int
	Name        string
	Description string
	Address     string
	Status      string
	Type        string

	Raw *gitlab.Runner
}
