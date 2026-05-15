package server

import "github.com/funtimecoding/go-library/pkg/tool/gohabiticad/face"

func New(c face.HabiticaSource) *Server {
	return &Server{habitica: c}
}
