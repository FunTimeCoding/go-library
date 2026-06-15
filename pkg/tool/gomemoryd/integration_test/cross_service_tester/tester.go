package cross_service_tester

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_client"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/integration_test/base"
)

type Tester struct {
	Goqueryd       *model_context_client.Client
	Gomemoryd      *model_context_client.Client
	GomemorydStore *store.Store
	goqueryd       *base.Server
}
