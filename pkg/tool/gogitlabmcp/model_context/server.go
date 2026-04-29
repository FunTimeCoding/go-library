package model_context

import (
	"github.com/getsentry/sentry-go"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type Server struct {
	client *gitlab.Client
	hub    *sentry.Hub
}
