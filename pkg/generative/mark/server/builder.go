package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

type Builder struct {
	name         string
	version      string
	instructions string
	logger       *logger.Logger
	verbose      bool
	resources    bool
	recorder     face.Recorder
}
