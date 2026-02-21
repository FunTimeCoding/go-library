package logger

import (
	"context"
	"log"
	"log/slog"
)

type Logger struct {
	context    context.Context
	structured *slog.Logger
	plain      *log.Logger
}
