package sentry

import "github.com/funtimecoding/go-library/pkg/errors/sentry/basic"

type Client struct {
	basic        *basic.Client
	organization string
	projects     []string
}
