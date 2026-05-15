package web_service_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/integration_test/base"
)

type Tester struct {
	server *base.Server
	Client *client.ClientWithResponses
}
