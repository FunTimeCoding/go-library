package goraidparsed

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"net/http"
	"testing"
)

func TestRunLifecycle(t *testing.T) {
	port, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				"",
				func(m *http.ServeMux) {
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New("", "", t.TempDir(), nil, nil),
							nil,
						),
						m,
					)
				},
			).WithListener(n),
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/status", base),
		http.StatusOK,
	)
}
