package model_context

import (
	"github.com/funtimecoding/go-library/pkg/iterm"
	"github.com/getsentry/sentry-go"
)

type Server struct {
	client *iterm.Client
	hub    *sentry.Hub
}
