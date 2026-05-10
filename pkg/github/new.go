package github

import (
	"context"
	"github.com/google/go-github/v86/github"
	"golang.org/x/oauth2"
)

func New(token string) *Client {
	result := &Client{context: context.Background()}
	result.client = github.NewClient(
		oauth2.NewClient(
			result.context,
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		),
	)

	return result
}
