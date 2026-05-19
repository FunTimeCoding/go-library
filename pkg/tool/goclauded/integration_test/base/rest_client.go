package base

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func (s *Server) RESTClient(t *testing.T) *client.ClientWithResponses {
	t.Helper()
	c, e := client.NewClientWithResponses(
		locator.New(
			constant.Localhost,
		).Insecure().Port(s.server.Port).String(),
	)
	assert.FatalOnError(t, e)

	return c
}
