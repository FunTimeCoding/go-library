package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"
)

type Server struct {
	store    *store.Store
	reporter face.Reporter
}
