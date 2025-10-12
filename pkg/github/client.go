package github

import (
	"context"
	"github.com/google/go-github/v75/github"
)

type Client struct {
	client  *github.Client
	context context.Context
}
