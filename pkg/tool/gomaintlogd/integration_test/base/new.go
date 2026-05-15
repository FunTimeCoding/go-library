package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.NewLite(filepath.Join(t.TempDir(), constant.TestDatabase))

	return &Server{
		Store: s,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(s, nil, constant.DefaultVersion).Mount(m)
				web.New(s).Mount(m)
			},
		),
	}
}
