package server

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/face"
)

type Server struct {
	habitica face.HabiticaSource
	reporter library.Reporter
}
