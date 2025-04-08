package runner

import "gitlab.com/gitlab-org/api/client-go"

type Runner struct {
	Identifier  int
	Name        string
	Description string
	Address     string
	Status      string
	Type        string
	Shared      bool
	Online      bool
	Paused      bool
	Tags        []string

	RawList   *gitlab.Runner
	RawDetail *gitlab.RunnerDetails
}
