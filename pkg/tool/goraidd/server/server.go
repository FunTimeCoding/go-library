package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
)

type Server struct {
	store      *store.Store
	outputPath string
	reporter   face.Reporter
}
