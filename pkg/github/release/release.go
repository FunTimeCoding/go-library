package release

import (
	"github.com/google/go-github/v59/github"
	"time"
)

type Release struct {
	Name      string
	CreatedAt time.Time

	Raw *github.RepositoryRelease
}
