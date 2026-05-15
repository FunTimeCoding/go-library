package worker_tester

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

type Tester struct {
	Store      *store.Store
	Worker     *worker.Worker
	MockClient *mock_client.Client
}
