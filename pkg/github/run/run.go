package run

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Run struct {
	Identifier int64
	Name       string
	Status     string
	CreatedAt  time.Time

	Raw *github.WorkflowRun
}
