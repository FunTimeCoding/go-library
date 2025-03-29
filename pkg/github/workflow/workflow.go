package workflow

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Workflow struct {
	Name      string
	CreatedAt time.Time

	Raw *github.Workflow
}
