package release

import (
	"github.com/google/go-github/v79/github"
	"time"
)

type Release struct {
	Name      string
	CreatedAt time.Time

	Raw *github.RepositoryRelease
}
