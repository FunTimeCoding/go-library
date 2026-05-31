package server

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/face"
)

func New(
	c face.HabiticaSource,
	r library.Reporter,
) *Server {
	return &Server{habitica: c, reporter: r}
}
