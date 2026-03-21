package release

import (
	"github.com/google/go-github/v84/github"
	"time"
)

type Release struct {
	Name   string
	Create time.Time
	Raw *github.RepositoryRelease
}
