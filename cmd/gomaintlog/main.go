package main

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mark3labs/mcp-go/server"
	"log"
)

func main() {
	a := constant.Listen
	log.Printf("Listen %s/mcp", a)
	s := web.Server(
		server.NewStreamableHTTPServer(
			maintenance_log.New().Nested(),
			server.WithStateful(true),
		),
		a,
	)
	web.ServeAsynchronous(s)
	system.KillSignalBlock()
	web.GracefulShutdown(context.Background(), s, false)
}
