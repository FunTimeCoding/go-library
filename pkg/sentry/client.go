package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/sentry/basic"
)

type Client struct {
	client       *sentry.Client
	basic        *basic.Client
	organization string
	projects     []string
}
