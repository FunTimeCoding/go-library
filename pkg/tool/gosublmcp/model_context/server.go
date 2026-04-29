package model_context

import (
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/getsentry/sentry-go"
)

type Server struct {
	client *sublime.Client
	hub    *sentry.Hub
}
