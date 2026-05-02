package recovery

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

type Recovery struct {
	logger   *logger.Logger
	reporter face.Reporter
}
