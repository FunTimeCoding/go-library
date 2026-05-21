package pull_request

import (
	"github.com/google/go-github/v87/github"
	"time"
)

type Request struct {
	Name   string
	Create time.Time
	Raw    *github.PullRequest
}
