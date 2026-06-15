package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/mock_indexer"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	i := mock_indexer.New()
	v := service.New(s, i, i, i)

	return &Server{
		t:       t,
		store:   s,
		indexer: i,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(v), m)
				model_context.New(
					v,
					memory.New(),
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(m)
			},
		),
	}
}
