package server

import "github.com/funtimecoding/go-library/pkg/habitica"

func New(c *habitica.Client) *Server {
	return &Server{habitica: c}
}
