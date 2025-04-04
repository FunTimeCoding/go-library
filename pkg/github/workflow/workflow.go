package workflow

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Workflow struct {
	Identifier int64
	Name       string
	State      string
	CreatedAt  time.Time

	Raw *github.Workflow
}
