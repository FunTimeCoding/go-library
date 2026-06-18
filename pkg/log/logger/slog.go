package logger

import "log/slog"

func (l *Logger) Slog() *slog.Logger {
	return l.structured
}
