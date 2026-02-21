package logger

import "github.com/funtimecoding/go-library/pkg/web/request_context"

func (l *Logger) Request(c *request_context.Context) {
	c.LogStart(l.context, l.structured)
}
