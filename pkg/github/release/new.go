package release

import "github.com/google/go-github/v70/github"

func New(v *github.RepositoryRelease) *Release {
	return &Release{Name: *v.TagName, CreatedAt: v.CreatedAt.Time, Raw: v}
}
