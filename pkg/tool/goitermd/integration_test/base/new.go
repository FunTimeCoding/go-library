package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/model_context"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			model_context.New(
				c,
				memory.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)

	return &Server{MockClient: c, ContextServer: v}
}
