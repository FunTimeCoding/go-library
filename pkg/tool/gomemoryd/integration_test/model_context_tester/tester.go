package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_client"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/base"
)

type Tester struct {
	*model_context_client.Client
	base *base.Server
}
