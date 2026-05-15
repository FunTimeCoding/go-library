package web_service_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/integration_test/base"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/mock_client"
)

type Tester struct {
	server     *base.Server
	Client     *client.ClientWithResponses
	MockClient *mock_client.Client
}
