package github

import (
	"context"
	"github.com/google/go-github/v78/github"
	"golang.org/x/oauth2"
)

func New(token string) *Client {
	client := &Client{context: context.Background()}
	client.client = github.NewClient(
		oauth2.NewClient(
			client.context,
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		),
	)

	return client
}
