package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) searchProjects(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	limit, f := r.RequireFloat(parameter.Limit)

	if f != nil {
		return response.Fail("limit is required: %v", f)
	}

	query := r.GetString(parameter.Query, "")
	includeDescriptions := r.GetBool(
		constant.IncludeDescriptions,
		false,
	)
	projects, h := s.jira.Projects()

	if h != nil {
		return s.captureFail(h, "Jira API unreachable")
	}

	var result []*server.JiraProject

	for _, p := range *projects {
		if len(result) >= int(limit) {
			break
		}

		if query != "" {
			keyMatch := strings.EqualFold(p.Key, query)
			nameMatch := strings.Contains(
				strings.ToLower(p.Name),
				strings.ToLower(query),
			)

			if !keyMatch && !nameMatch {
				continue
			}
		}

		if includeDescriptions {
			detail, i := s.jira.Project(p.Key)

			if i != nil {
				return s.captureFail(i, "project not found")
			}

			result = append(result, convert.JiraProject(detail))
		} else {
			o := &server.JiraProject{Key: p.Key, Name: p.Name}

			if p.ProjectTypeKey != "" {
				o.Type = &p.ProjectTypeKey
			}

			result = append(result, o)
		}
	}

	return response.SuccessAny(result)
}
