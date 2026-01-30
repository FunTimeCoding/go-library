package github

import (
	"context"
	"github.com/google/go-github/v82/github"
)

type Client struct {
	client  *github.Client
	context context.Context
}
