package main

import (
	"context"
	mcpServer "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func main() {
	s := mcpServer.New(maintenance_log.New().Nested())
	m := http.NewServeMux()
	s.Setup(m)
	h := web.Server(m, mcpServer.Address)
	web.ServeAsynchronous(h)
	system.KillSignalBlock()
	web.GracefulShutdown(context.Background(), h, true)
}
