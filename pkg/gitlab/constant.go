package gitlab

import "gitlab.com/gitlab-org/api/client-go"

const (
	OpenedState  = "opened"
	PendingState = "pending"
)

// Identifier Sort field
const Identifier = "id"

// Sort order
const (
	Ascending  = "asc"
	Descending = "desc"
)

// Job status
const (
	Success = "success"
	Failed  = "failed"
)

var DefaultListOptions = gitlab.ListOptions{Page: 0, PerPage: 100}

const (
	HostEnvironment  = "GITLAB_HOST"
	TokenEnvironment = "GITLAB_TOKEN"

	PrivateTokenHeader = "Private-Token"
	JobTokenHeader     = "Job-Token"

	LatestVersion = "latest"

	PerPage1000 int = 1000
	PerPage100  int = 100
)

// Environment variables during jobs
const (
	InterfaceLocator  = "CI_API_V4_URL"
	ProjectIdentifier = "CI_PROJECT_ID"
	CommitTag         = "CI_COMMIT_TAG"
	JobToken          = "CI_JOB_TOKEN"
)
