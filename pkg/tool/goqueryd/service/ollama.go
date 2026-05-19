package service

import "github.com/funtimecoding/go-library/pkg/generative/ollama"

func (s *Service) Ollama() *ollama.Client {
	return s.ollama
}
