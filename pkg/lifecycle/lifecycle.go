package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

type Lifecycle struct {
	component []face.Component
	logger    *logger.Logger
	verbose   bool
}
