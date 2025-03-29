package job

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Job struct {
	Identifier int64
	Name       string
	Hash       string
	CreatedAt  time.Time

	Raw *github.WorkflowJob
}
