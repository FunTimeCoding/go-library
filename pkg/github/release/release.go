package release

import (
	"github.com/google/go-github/v82/github"
	"time"
)

type Release struct {
	Name   string
	Create time.Time

	Raw *github.RepositoryRelease
}
