package issue

import (
	"github.com/google/go-github/v59/github"
	"net/url"
	"strings"
)

func New(input *github.Issue) *Issue {
	var repository string

	if u, e := url.Parse(*input.RepositoryURL); e == nil {
		repository = strings.TrimPrefix(u.Path, "/repos/")
	}

	return &Issue{
		Repository: repository,
		Title:      *input.Title,
		Link:       *input.HTMLURL,
		Raw:        input,
	}
}
