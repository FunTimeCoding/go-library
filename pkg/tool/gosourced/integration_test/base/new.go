package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"
	"github.com/funtimecoding/go-library/pkg/tool/gosourced/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosourced/service"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	fixture string,
) *Server {
	t.Helper()
	d := testutil.PrepareTestPackage(t, fixture)
	i := inventory.New()
	i.Modules = []inventory.Module{
		{Name: "test", Directory: d},
	}
	s := service.New(i)
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			model_context.New(
				s,
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)

	return &Server{ContextServer: v, Directory: d}
}
