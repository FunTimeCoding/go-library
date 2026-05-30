package request

import (
	"context"
	"errors"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/mark3labs/mcp-go/util"
)

func Hooks(
	h *server.Hooks,
	l util.Logger,
	verbose bool,
) {
	h.AddBeforeAny(
		func(
			_ context.Context,
			identifier any,
			m mcp.MCPMethod,
			a any,
		) {
			if q, okay := a.(*mcp.InitializeRequest); okay {
				l.Infof("%s\n", q.Method)

				return
			}

			if q, okay := a.(*mcp.ListToolsRequest); okay {
				l.Infof("%s\n", q.Method)

				return
			}

			if q, okay := a.(*mcp.CallToolRequest); okay {
				l.Infof(
					"%s %s %+v\n",
					q.Method,
					q.Params.Name,
					q.Params.Arguments,
				)

				return
			}

			if q, okay := a.(*mcp.PaginatedRequest); okay {
				l.Infof("PaginatedRequest: %s\n", q.Method)
				LogParameter(l, &q.Request)

				return
			}

			if q, okay := a.(*mcp.Request); okay {
				l.Infof("Request: %s\n", q.Method)
				LogParameter(l, q)

				return
			}

			l.Errorf(
				"Unexpected type: %T %+v %s %+v\n",
				a,
				identifier,
				m,
				a,
			)
		},
	)

	if verbose {
		h.AddBeforeInitialize(
			func(
				_ context.Context,
				identifier any,
				q *mcp.InitializeRequest,
			) {
				l.Infof("Before initialize: %+v %+v\n", identifier, q)
			},
		)
		h.AddOnRegisterSession(
			func(
				_ context.Context,
				s server.ClientSession,
			) {
				l.Infof("Session registered: %v\n", s.SessionID())
			},
		)
		h.AddOnRequestInitialization(
			func(
				_ context.Context,
				identifier any,
				a any,
			) error {
				if q, okay := a.(*mcp.ListToolsRequest); okay {
					if false {
						l.Infof(
							"ListToolsRequest: %+v %+v\n",
							identifier,
							q,
						)

						return errors.New("unauthorized")
					}
				}

				return nil
			},
		)
	}
}
