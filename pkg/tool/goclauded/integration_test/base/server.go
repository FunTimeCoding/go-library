package base

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
)

type Server struct {
	*service_tester.Tester
	server *model_context_server.Server
}
