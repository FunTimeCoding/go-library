package run

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Run struct {
	Name      string
	CreatedAt time.Time

	Raw *github.WorkflowRun
}
