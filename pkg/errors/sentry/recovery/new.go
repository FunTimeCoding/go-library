package recovery

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

func New(
	l *logger.Logger,
	r face.Reporter,
) *Recovery {
	return &Recovery{logger: l, reporter: r}
}
