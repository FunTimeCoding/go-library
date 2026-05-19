package base

import "github.com/funtimecoding/go-library/pkg/generative/ollama"

func (s *Server) Ollama() *ollama.Client {
	return s.ollama
}
