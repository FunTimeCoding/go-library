package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	goqueryd "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	generated "github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	l := ollama.NewEnvironment()
	a, e := rerank.New(
		environment.Required(goqueryd.ModelEnvironment),
		environment.Required(goqueryd.TokenizerEnvironment),
	)
	errors.PanicOnError(e)
	t.Cleanup(func() { errors.LogClose(a) })
	v := service.New(s, l, a)

	return &Server{
		t:        t,
		store:    s,
		ollama:   l,
		reranker: a,
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
