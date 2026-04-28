package goraidparsed

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"net/http"
	"testing"
)

func TestRunLifecycle(t *testing.T) {
	port, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithListener(
			n,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New("", "", t.TempDir(), nil, nil),
					m,
				)
			},
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
