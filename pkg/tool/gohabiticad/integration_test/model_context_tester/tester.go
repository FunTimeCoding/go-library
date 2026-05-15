package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_client"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/integration_test/base"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/mock_client"
)

type Tester struct {
	server     *base.Server
	Client     *model_context_client.Client
	MockClient *mock_client.Client
}
