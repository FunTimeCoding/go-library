package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	generated "github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/server"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			generated.HandlerFromMux(server.New(c), m)
			model_context.New(
				c,
				memory.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)

	return &Server{MockClient: c, ContextServer: v}
}
