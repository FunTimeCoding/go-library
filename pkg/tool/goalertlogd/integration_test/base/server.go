package base

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

type Server struct {
	Store         *store.Store
	Worker        *worker.Worker
	MockClient    *mock_client.Client
	ContextServer *model_context_server.Server
}
