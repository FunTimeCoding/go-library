package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) listInstances(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListInstances,
) (*mcp.CallToolResult, error) {
	var active string

	if session := server.ClientSessionFromContext(x); session != nil {
		active, _ = s.service.ActiveInstance(session.SessionID())
	}

	type entry struct {
		Name             string `json:"name"`
		AlertmanagerHost string `json:"alertmanager_host"`
		PrometheusHost   string `json:"prometheus_host"`
		Active           bool   `json:"active"`
	}
	var result []entry

	for _, i := range s.service.Instances() {
		result = append(
			result,
			entry{
				Name:             i.Name,
				AlertmanagerHost: i.Alertmanager.Host,
				PrometheusHost:   i.Prometheus.Host,
				Active:           i.Name == active,
			},
		)
	}

	return response.SuccessAny(result)
}
