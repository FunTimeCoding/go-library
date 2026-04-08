package gomaintlogd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"io"
	"net/http"
	"path/filepath"
	"testing"
)

func setup(t *testing.T) (*store.Store, int, *lifecycle.Lifecycle) {
	t.Helper()
	s := store.NewSQLite(filepath.Join(t.TempDir(), "test.db"))
	port := system.FindUnusedPort(19600)
	l := lifecycle.New(
		lifecycle.WithServerTimeout(
			fmt.Sprintf(":%d", port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
				web.NewServer(s).Mount(m)
			},
			0,
			0,
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
