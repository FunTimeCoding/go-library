package job

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Job struct {
	Name      string
	CreatedAt time.Time

	Raw *github.WorkflowJob
}
