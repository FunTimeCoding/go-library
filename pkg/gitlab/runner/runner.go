package runner

import "gitlab.com/gitlab-org/api/client-go"

type Runner struct {
	Identifier  int
	Name        string
	Description string
	Status      string
	Type        string
	Shared      bool
	Online      bool
	Paused      bool
	Tags        []string

	Address string // Loaded separately via GraphQL

	concern []string

	RawList   *gitlab.Runner
	RawDetail *gitlab.RunnerDetails
}
