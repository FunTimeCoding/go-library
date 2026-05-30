package adapter

import "github.com/funtimecoding/go-library/pkg/log/logger"

func New(l *logger.Logger) *Adapter {
	return &Adapter{Logger: l}
}
