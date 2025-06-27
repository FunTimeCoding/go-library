package run

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Run struct {
	MonitorIdentifier string
	Identifier        int64
	Name              string
	Status            string
	Create            time.Time
	Update            time.Time

	Raw *github.WorkflowRun
}
