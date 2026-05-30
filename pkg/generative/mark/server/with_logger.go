package server

import "github.com/funtimecoding/go-library/pkg/log/logger"

func (b *Builder) WithLogger(
	l *logger.Logger,
	verbose bool,
) *Builder {
	b.logger = l
	b.verbose = verbose

	return b
}
