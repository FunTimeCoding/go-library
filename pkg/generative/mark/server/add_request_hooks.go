package server

import (
	"github.com/funtimecoding/go-library/pkg/generative/mark/adapter"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/mark3labs/mcp-go/server"
)

func addRequestHooks(
	h *server.Hooks,
	l *logger.Logger,
	verbose bool,
) {
	request.Hooks(h, adapter.New(l), verbose)
}
