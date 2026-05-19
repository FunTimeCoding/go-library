package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/notifier"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/mock_indexer"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()
	now := time.Now()
	s := store.New(
		filepath.Join(t.TempDir(), constant.TestDatabase),
		notifier.New(),
		func() time.Time { return now },
	)
	l := logger.New(t.Context())
	i := mock_indexer.New()
	v := service.New(s, mock_client.New())

	return &Server{
		t:       t,
		store:   s,
		indexer: i,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					server.New(v, nil, l, t.TempDir(), t.TempDir()),
					m,
				)
				model_context.New(
					v,
					i,
					memory.New(),
					l,
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(m)
			},
		),
		now: &now,
	}
}
