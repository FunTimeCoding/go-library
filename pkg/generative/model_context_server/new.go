package model_context_server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	setup func(*http.ServeMux),
) *Server {
	t.Helper()
	p, n := system.ClaimPort()
	b := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithListener(n, setup),
	)
	b.Run()
	assert.Listen(t, p)

	return &Server{Port: p, Lifecycle: b}
}
