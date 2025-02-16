package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/sentry/basic_client"
)

type Client struct {
	client       *sentry.Client
	basicClient  *basic_client.Client
	organization string
	projects     []string
}
