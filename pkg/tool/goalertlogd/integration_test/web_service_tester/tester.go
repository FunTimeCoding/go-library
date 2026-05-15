package web_service_tester

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/integration_test/base"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

type Tester struct {
	server     *base.Server
	Client     *client.ClientWithResponses
	Worker     *worker.Worker
	MockClient *mock_client.Client
}
