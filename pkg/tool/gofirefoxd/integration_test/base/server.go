package base

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/mock_client"
)

type Server struct {
	MockClient    *mock_client.Client
	ContextServer *model_context_server.Server
}
