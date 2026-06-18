package request

import (
	"github.com/mark3labs/mcp-go/mcp"
	"log/slog"
)

func LogParameter(
	l *slog.Logger,
	q *mcp.Request,
) {
	if q.Params.Meta != nil {
		if len(q.Params.Meta.AdditionalFields) > 0 {
			for k, v := range q.Params.Meta.AdditionalFields {
				l.Info("parameter", k, v)
			}
		} else {
			l.Info("Empty parameters")
		}
	} else {
		l.Info("No parameters")
	}
}
