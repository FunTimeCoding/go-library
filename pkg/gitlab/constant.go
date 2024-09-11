package gitlab

import "github.com/xanzy/go-gitlab"

const (
	OpenedState  = "opened"
	PendingState = "pending"
)

var DefaultListOptions = gitlab.ListOptions{Page: 0, PerPage: 100}

const (
	Host  = "GITLAB_HOST"
	Token = "GITLAB_TOKEN"

	PrivateTokenHeader = "Private-Token"
	JobTokenHeader     = "Job-Token"

	LatestVersion = "latest"

	PerPage = int(1000)
)

// Environment variables during jobs
const (
	InterfaceLocator  = "CI_API_V4_URL"
	ProjectIdentifier = "CI_PROJECT_ID"
	CommitTag         = "CI_COMMIT_TAG"
	JobToken          = "CI_JOB_TOKEN"
)
