package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"gitlab.com/gitlab-org/api/client-go"
)

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
	GroupEnvironment = "GITLAB_GROUP"

	PrivateTokenHeader = "Private-Token"
	JobTokenHeader     = "Job-Token"

	LatestVersion = "latest"

	LatestSuffix = ":latest"

	PerPage1000 int = 1000
	PerPage100  int = 100
)

// Environment variables during jobs
const (
	CommitTag         = "CI_COMMIT_TAG"
	InterfaceLocator  = "CI_API_V4_URL"
	JobToken          = "CI_JOB_TOKEN"
	ProjectIdentifier = "CI_PROJECT_ID"
	ProjectNamespace  = "CI_PROJECT_NAMESPACE"
	QualifiedName     = "CI_SERVER_FQDN"
)

var (
	Format      = option.Color.Copy()
	CheckFormat = Format.Copy().Tag(tag.Project)
)
