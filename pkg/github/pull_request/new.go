package pull_request

import "github.com/google/go-github/v82/github"

func New(v *github.PullRequest) *Request {
	return &Request{
		Name:   v.GetTitle(),
		Create: v.GetCreatedAt().Time,
		Raw:    v,
	}
}
