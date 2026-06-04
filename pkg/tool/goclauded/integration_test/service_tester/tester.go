package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/mock_indexer"
)

type Tester struct {
	*store_tester.Tester
	Service  *service.Service
	Client   *mock_client.Client
	Indexer  *mock_indexer.Indexer
	Notifier *mock_notifier.Notifier
	Harbor   string
}
