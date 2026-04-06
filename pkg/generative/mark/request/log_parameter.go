package request

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/util"
)

func LogParameter(
	l util.Logger,
	q *mcp.Request,
) {
	if q.Params.Meta != nil {
		if len(q.Params.Meta.AdditionalFields) > 0 {
			for k, v := range q.Params.Meta.AdditionalFields {
				l.Infof("  %s=%+v\n", k, v)
			}
		} else {
			l.Infof("  Empty parameters")
		}
	} else {
		l.Infof("  No parameters")
	}
}
