package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/mock_service"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/mock_snippet"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	c.AddNode("test")
	svc := mock_service.New("test", c, mock_snippet.New())
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			model_context.New(
				svc,
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)

	return &Server{MockClient: c, ContextServer: v}
}
