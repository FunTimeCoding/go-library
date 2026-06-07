package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
)

type Server struct {
	service           *service.Service
	logger            *logger.Logger
	reporter          face.Reporter
	harborPath        string
	sessionExportPath string
}
