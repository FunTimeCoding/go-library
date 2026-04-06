package request

import (
	"context"
	"errors"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/mark3labs/mcp-go/util"
)

func Hooks(
	l util.Logger,
	verbose bool,
) *server.Hooks {
	result := &server.Hooks{}
	result.AddBeforeAny(
		func(
			_ context.Context,
			identifier any,
			m mcp.MCPMethod,
			a any,
		) {
			if q, ok := a.(*mcp.InitializeRequest); ok {
				l.Infof("%s\n", q.Method)

				return
			}

			if q, ok := a.(*mcp.ListToolsRequest); ok {
				l.Infof("%s\n", q.Method)

				return
			}

			if q, ok := a.(*mcp.CallToolRequest); ok {
				l.Infof(
					"%s %s %+v\n",
					q.Method,
					q.Params.Name,
					q.Params.Arguments,
				)

				return
			}

			if q, ok := a.(*mcp.PaginatedRequest); ok {
				l.Infof("PaginatedRequest: %s\n", q.Method)
				LogParameter(l, &q.Request)

				return
			}

			if q, ok := a.(*mcp.Request); ok {
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
		result.AddBeforeInitialize(
			func(
				_ context.Context,
				identifier any,
				q *mcp.InitializeRequest,
			) {
				l.Infof("Before initialize: %+v %+v\n", identifier, q)
			},
		)
		result.AddOnRegisterSession(
			func(
				_ context.Context,
				s server.ClientSession,
			) {
				l.Infof("Session registered: %v\n", s.SessionID())
			},
		)
		result.AddOnRequestInitialization(
			func(
				_ context.Context,
				identifier any,
				a any,
			) error {
				if q, ok := a.(*mcp.ListToolsRequest); ok {
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

	return result
}
