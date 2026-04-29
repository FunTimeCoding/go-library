package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"io"
	"net/http"
	"path/filepath"
	"testing"
)

func setup(t *testing.T) (*store.Store, int, *lifecycle.Lifecycle) {
	t.Helper()
	s := store.NewLite(filepath.Join(t.TempDir(), "test.db"))
	port, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithListener(
			n,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
				web.New(s).Mount(m)
			},
		),
	)
	l.Run()
	assert.Listen(t, port)

	return s, port, l
}

func getBody(
	t *testing.T,
	locator string,
) string {
	t.Helper()
	r, e := http.Get(locator)
	assert.FatalOnError(t, e)

	return readBody(t, r)
}

func readBody(
	t *testing.T,
	r *http.Response,
) string {
	t.Helper()
	defer errors.PanicClose(r.Body)
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)

	return string(b)
}
