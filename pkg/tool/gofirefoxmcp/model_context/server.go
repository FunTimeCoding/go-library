package model_context

import (
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/getsentry/sentry-go"
)

type Server struct {
	client *firefox.Client
	hub    *sentry.Hub
}
