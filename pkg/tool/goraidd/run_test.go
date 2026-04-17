package goraidd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/web"
	"net/http"
	"testing"
)

func TestRunLifecycle(t *testing.T) {
	logCachePath := t.TempDir()
	outputPath := t.TempDir()
	port := system.FindUnusedPort(19600)
	address := fmt.Sprintf(":%d", port)
	l := lifecycle.New(
		lifecycle.WithServer(
			address,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New(logCachePath, outputPath),
					m,
				)
				web.NewServer(logCachePath, "", outputPath).Mount(m)
			},
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/reports", base),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/reports", base),
		http.StatusOK,
	)
}
