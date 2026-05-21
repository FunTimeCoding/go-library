package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/server"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := service_tester.New(t)
	l := logger.New(t.Context())

	return &Server{
		Tester: s,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					server.New(s.Service, l, t.TempDir(), t.TempDir()),
					m,
				)
				model_context.New(
					s.Service,
					memory.New(),
					l,
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(m)
			},
		),
	}
}
