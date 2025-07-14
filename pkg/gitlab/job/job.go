package job

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Job struct {
	MonitorIdentifier string
	Identifier        int
	Name              string
	Status            string
	Stage             string
	Create            *time.Time
	Link              string

	Project *project.Project
	concern []string

	Raw *gitlab.Job
}
