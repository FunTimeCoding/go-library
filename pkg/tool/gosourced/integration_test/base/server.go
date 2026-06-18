package base

import "github.com/funtimecoding/go-library/pkg/generative/model_context_server"

type Server struct {
	ContextServer *model_context_server.Server
	Directory     string
}
