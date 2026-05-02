package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

type Server struct {
	parserPath   string
	templatePath string
	outputPath   string
	logger       *logger.Logger
	reporter     face.Reporter
}
