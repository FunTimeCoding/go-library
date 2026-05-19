package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/integration_test/base"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	t.Cleanup(s.Close)

	return &Tester{
		Service: service.New(s.Store(), s.Ollama(), s.Reranker()),
	}
}
