package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/mock_indexer"
)

type Tester struct {
	Service *service.Service
	Indexer *mock_indexer.Indexer
}
