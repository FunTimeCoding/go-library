package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_client"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	t.Cleanup(s.Close)
	c := model_context_client.New(t, s.Port())
	t.Cleanup(c.Close)

	return &Tester{
		Client: c,
		base:   s,
	}
}
