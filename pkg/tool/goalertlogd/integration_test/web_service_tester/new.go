package web_service_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/integration_test/base"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	c, e := client.NewClientWithResponses(
		locator.New(
			constant.Localhost,
		).Insecure().Port(s.ContextServer.Port).String(),
	)
	assert.FatalOnError(t, e)

	return &Tester{
		server:     s,
		Client:     c,
		Worker:     s.Worker,
		MockClient: s.MockClient,
	}
}
