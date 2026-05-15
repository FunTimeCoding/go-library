package base

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
)

type Server struct {
	Store         *store.Store
	ContextServer *model_context_server.Server
}
