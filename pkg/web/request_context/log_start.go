package request_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"log/slog"
)

func (c *Context) LogStart(
	x context.Context,
	l *slog.Logger,
) {
	l.LogAttrs(x, slog.LevelInfo, constant.RequestStart, c.Attribute()...)
}
