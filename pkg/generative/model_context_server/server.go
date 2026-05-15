package model_context_server

import "github.com/funtimecoding/go-library/pkg/lifecycle"

type Server struct {
	Port      int
	Lifecycle *lifecycle.Lifecycle
}
