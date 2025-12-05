package job

import (
	"github.com/google/go-github/v80/github"
	"time"
)

type Job struct {
	Identifier int64
	Name       string
	Status     string
	Conclusion string
	Hash       string
	CreatedAt  time.Time

	Raw *github.WorkflowJob
}
