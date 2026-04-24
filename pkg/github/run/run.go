package run

import (
	"github.com/funtimecoding/go-library/pkg/github/job"
	"github.com/funtimecoding/go-library/pkg/github/repository"
	"github.com/google/go-github/v85/github"
	"time"
)

type Run struct {
	MonitorIdentifier string
	Identifier        int64
	Name              string
	Status            string
	Conclusion        string
	Branch            string
	Create            time.Time
	Update            time.Time
	concern           []string
	Repository        *repository.Repository
	Jobs              []*job.Job
	Raw               *github.WorkflowRun
}
