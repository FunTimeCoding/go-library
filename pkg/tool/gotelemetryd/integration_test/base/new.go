package base

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	generated "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	database, e := gorm.Open(
		sqlite.Open(filepath.Join(t.TempDir(), constant.TestDatabase)),
		&gorm.Config{},
	)
	errors.PanicOnError(e)
	s := store.New(database)
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			generated.HandlerFromMux(server.New(s), m)
			model_context.New(
				s,
				memory.New(),
				constant.DefaultVersion,
			).Mount(m)
			web.New(s).Mount(m)
		},
	)

	return &Server{Store: s, ContextServer: v}
}
