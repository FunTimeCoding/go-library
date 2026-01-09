package release

import "github.com/google/go-github/v81/github"

func New(v *github.RepositoryRelease) *Release {
	return &Release{Name: *v.TagName, Create: v.CreatedAt.Time, Raw: v}
}
