package service_tester

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	memoryMock "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store_tester.New(t)
	c := mock_client.New()
	i := mock_indexer.New()
	n := mock_notifier.New()
	harbor := t.TempDir()

	return &Tester{
		Store: s,
		t:     t,
		Service: service.New(
			s.Store,
			c,
			memoryMock.New(),
			i,
			n,
			reporter.NewOptional("", ""),
			harbor,
			s.Clock(),
			logger.New(context.Background()),
		),
		Harbor:   harbor,
		Client:   c,
		Indexer:  i,
		Notifier: n,
	}
}
